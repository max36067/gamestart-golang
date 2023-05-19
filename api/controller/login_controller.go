package controller

import (
	"apigee-portal/v2/bootstrap"
	"apigee-portal/v2/domain"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginUsecase       domain.LoginUsecase
	SaltRepository     domain.SaltRepository
	GoogleOauthUsecase domain.GoogleOauthUsecase
	Env                *bootstrap.Env
}

func (lc LoginController) ServerLogin(c *gin.Context) {
	var request domain.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := lc.LoginUsecase.GetUserByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found with the given email."})
		return
	}

	saltString, err := lc.SaltRepository.GetSaltByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Invalid credentials."})
		return
	}
	request.Password = request.Password + saltString

	if lc.LoginUsecase.VerifyPassword(user.Password, request.Password) != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Invalid credentials."})
		return
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.SecretKey, lc.Env.ExpiredMinutes)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&user, lc.Env.SecretKey, lc.Env.ExpiredMinutes)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}

func (lc *LoginController) GoogleOauth(c *gin.Context) {
	redirectUri := "http://localhost:8080/api/v1/oauth/google/login"
	queryParams := url.Values{}
	queryParams.Set("client_id", lc.Env.GoogleOauthClientID)
	queryParams.Set("response_type", "code")
	queryParams.Set("scope", "https://www.googleapis.com/auth/userinfo.profile")
	queryParams.Set("redirect_uri", redirectUri)
	queryParams.Set("access_type", "offline")
	oauthUrl := fmt.Sprintf("%s?%s", lc.Env.GoogleOauthAuthUri, queryParams.Encode())
	c.JSON(http.StatusOK, gin.H{"google_oauth_url": oauthUrl})
}

func (lc *LoginController) GoogleOauthLogin(c *gin.Context) {
	code := c.Query("code")
	token, err := lc.GoogleOauthUsecase.RequestAccessToken(code)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	userInfo, err := lc.GoogleOauthUsecase.GetUserInfo(&token)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"message": "Can not get user info from Google."})
		return
	}

	c.JSON(http.StatusOK, userInfo)
}

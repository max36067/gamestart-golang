package controller

import (
	"apigee-portal/v2/bootstrap"
	"apigee-portal/v2/domain"
	"apigee-portal/v2/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *bootstrap.Env
}

// @Summary User Signup
// @Description User Signup
// @Tags Signup
// @Accept json
// @Produce json
// @Param user_info body domain.SignupRequest false "userInfo"
// @Success 200 {object} domain.SignupResponse
// @Router /signup [post]
func (sc *SignupController) Register(c *gin.Context) {
	var request domain.SignupRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	_, err := sc.SignupUsecase.GetUserByEmail(request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "User is already exists."})
		return
	}

	saltString, err := utils.GenerateSalt()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(request.Password + saltString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	salt := domain.Salt{
		Email: request.Email,
		Salt:  saltString,
	}
	err = sc.SignupUsecase.CreateSalt(&salt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user := domain.User{
		Email:       request.Email,
		Password:    hashedPassword,
		Name:        request.Name,
		IsActive:    true,
		IsSuperUser: false,
	}
	err = sc.SignupUsecase.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	accessToken, err := sc.SignupUsecase.CreateAccessToken(&user, sc.Env.SecretKey, sc.Env.ExpiredMinutes)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(&user, sc.Env.SecretKey, sc.Env.ExpiredMinutes)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)

}

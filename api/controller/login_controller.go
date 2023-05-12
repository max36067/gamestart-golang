package controller

import (
	"apigee-portal/v2/bootstrap"
	"apigee-portal/v2/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginUsecase   domain.LoginUsecase
	SaltRepository domain.SaltRepository
	Env            *bootstrap.Env
}

func (lc LoginController) Login(c *gin.Context) {
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

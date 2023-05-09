package controller

import (
	"apigee-portal/v2/bootstrap"
	"apigee-portal/v2/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserRepository domain.UserRepository
	Env            *bootstrap.Env
}

func (uc *UserController) List(c *gin.Context) {
	users, err := uc.UserRepository.Fetch()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

package controllers

import (
	"apigee-portal/v2/models"
	"net/http"

	"apigee-portal/v2/repository"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var user []models.UserResponse
	repository.Session.Find(&user)
	c.JSON(http.StatusOK, gin.H{
		"users": user,
	})
}

func GetUser(c *gin.Context) {
	var user models.User
	var requestBody models.UserRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := repository.Session.Where("email = ?", requestBody.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	userResponse := &models.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(200, gin.H{"user": userResponse})
}

func CreateUser(c *gin.Context) {

}

package controllers

import (
	"apigee-portal/v2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	var payload *models.SignInInput
	if err := c.ShouldBindJSON(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
	}

	// var user models.User
	// result := models.Session.First(&user, "email = ?", payload.Email)

}

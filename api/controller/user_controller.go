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

// @Summary User List
// @Description User List
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.UserList
// @Router /users [get]
func (uc *UserController) List(c *gin.Context) {
	users, err := uc.UserRepository.Fetch()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, domain.UserList{Users: users})
}

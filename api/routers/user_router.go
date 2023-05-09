package router

import (
	"apigee-portal/v2/api/controller"
	"apigee-portal/v2/bootstrap"
	"apigee-portal/v2/postgres"
	"apigee-portal/v2/repository"

	"github.com/gin-gonic/gin"
)

func NewUserRouter(db postgres.DataBase, env *bootstrap.Env, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	uc := controller.UserController{
		UserRepository: ur,
		Env:            env,
	}
	group.GET("/users", uc.List)
}

package router

import (
	"apigee-portal/v2/api/controller"
	"apigee-portal/v2/bootstrap"
	"apigee-portal/v2/repository"
	"apigee-portal/v2/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewLogoutRoute(databasses *bootstrap.Databases, timeout time.Duration, env *bootstrap.Env, group *gin.RouterGroup) {
	lp := repository.NewTokenBlacklistRepository(databasses.RDB)
	lu := usecase.NewLogoutUsecase(lp, timeout)
	lc := controller.LogoutController{
		LogoutUsecase: lu,
		Env:           env,
	}

	group.GET("/logout", lc.Logout)
}

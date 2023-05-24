package router

import (
	"apigee-portal/v2/api/controller"
	"apigee-portal/v2/bootstrap"
	"apigee-portal/v2/repository"
	"apigee-portal/v2/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRefreshTokenRoute(db *gorm.DB, timeout time.Duration, env *bootstrap.Env, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	refreshTokenController := controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}

	group.POST("/refresh", refreshTokenController.RefreshToken)
}

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

func NewSignupRouter(db *gorm.DB, timeout time.Duration, env *bootstrap.Env, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	sr := repository.NewSaltRepository(db)
	su := usecase.NewSignupUsecase(ur, sr, timeout)
	signupController := controller.SignupController{
		SignupUsecase: su,
		Env:           env,
	}

	group.POST("/signup", signupController.Register)
}

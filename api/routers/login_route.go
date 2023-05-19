package router

import (
	"apigee-portal/v2/api/controller"
	"apigee-portal/v2/bootstrap"
	"apigee-portal/v2/repository"
	"apigee-portal/v2/usecase"
	"apigee-portal/v2/usecase/google_usecase"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewLoginRoute(db *gorm.DB, timeout time.Duration, env *bootstrap.Env, group *gin.RouterGroup) {
	sr := repository.NewSaltRepository(db)
	ur := repository.NewUserRepository(db)
	gou := google_usecase.NewGoogleOauthUsecase(env)
	lu := usecase.NewLoginUsecase(ur, timeout)
	loginController := controller.LoginController{
		LoginUsecase:       lu,
		SaltRepository:     sr,
		GoogleOauthUsecase: gou,
		Env:                env,
	}

	group.POST("/login", loginController.ServerLogin)

	oauth := group.Group("/oauth")
	oauth.GET("/google/url", loginController.GoogleOauth)
	oauth.GET("/google/login", loginController.GoogleOauthLogin)
}

package router

import (
	"apigee-portal/v2/api/middleware"
	"apigee-portal/v2/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
	apiV1Router := gin.Group("/api/v1")
	publicRouter := apiV1Router.Group("")
	NewLoginRoute(db, timeout, env, publicRouter)

	protectedRouter := apiV1Router.Group("")
	protectedRouter.Use(middleware.JWTAuthMiddleware(env.SecretKey))
	NewUserRouter(db, env, protectedRouter)
	NewRefreshTokenRoute(db, timeout, env, protectedRouter)
}

package router

import (
	"apigee-portal/v2/api/middleware"
	"apigee-portal/v2/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, databases *bootstrap.Databases, gin *gin.Engine) {
	apiV1Router := gin.Group("/api/v1")
	publicRouter := apiV1Router.Group("")
	NewLoginRoute(databases.DB, timeout, env, publicRouter)
	NewSignupRouter(databases.DB, timeout, env, publicRouter)

	protectedRouter := apiV1Router.Group("")
	protectedRouter.Use(middleware.JWTAuthMiddleware(databases.RDB, env.SecretKey))
	NewUserRouter(databases.DB, env, protectedRouter)
	NewRefreshTokenRoute(databases.DB, timeout, env, protectedRouter)
	NewLogoutRoute(databases, timeout, env, protectedRouter)
}

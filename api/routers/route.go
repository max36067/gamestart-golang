package router

import (
	"apigee-portal/v2/bootstrap"
	"apigee-portal/v2/postgres"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, db postgres.DataBase, gin *gin.Engine) {
	// publicRouter := gin.Group("")

	protectedRouter := gin.Group("/api/v1")
	NewUserRouter(db, env, protectedRouter)
}

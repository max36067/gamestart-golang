package routers

import (
	"apigee-portal/v2/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {

	router.GET("/users", controllers.GetUsers)
	router.POST("/user", controllers.GetUser)
}

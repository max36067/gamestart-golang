package routers

import (
	"apigee-portal/v2/routers/v1"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	routers.UserRouter(v1)
}

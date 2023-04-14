package main

import (
	"apigee-portal/v2/routers"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.New()
	routers.RegisterRouter(router)
}

func main() {
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

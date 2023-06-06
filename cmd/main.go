package main

import (
	router "apigee-portal/v2/api/routers"
	"apigee-portal/v2/bootstrap"
	"apigee-portal/v2/docs"
	"time"

	"github.com/gin-contrib/cors"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger

// @contact.name Max.Huang

// @host localhost:8080
// @securitydefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @schemes http
func main() {
	app := bootstrap.NewApp()
	env := app.Env

	server := gin.New()
	timeout := time.Duration(time.Second * time.Duration(env.SystemTimeout))
	// cors
	server.Use(cors.New(app.Cors))

	// swagger setting
	docs.SwaggerInfo.BasePath = "/api/v1"
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// default server setup
	router.Setup(env, timeout, app.Databases, server)
	server.GET("/home", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hollow World"})
	})
	server.Run(":8080")
}

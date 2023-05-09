package main

import (
	router "apigee-portal/v2/api/routers"
	"apigee-portal/v2/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.NewApp()
	env := app.Env

	server := gin.New()

	router.Setup(env, app.Postgresql, server)

	server.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

package main

import (
	router "apigee-portal/v2/api/routers"
	"apigee-portal/v2/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.NewApp()
	env := app.Env

	server := gin.New()
	timeout := time.Duration(time.Second * time.Duration(env.SystemTimeout))

	router.Setup(env, timeout, app.Database, server)

	server.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

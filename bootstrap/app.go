package bootstrap

import (
	"apigee-portal/v2/postgres"
	"log"

	"github.com/joho/godotenv"
)

type Application struct {
	Env        *Env
	Postgresql postgres.DataBase
}

func NewApp() *Application {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := &Application{}
	app.Env = NewEnv()
	app.Postgresql = NewPostgresDatabase(app.Env)

	return app
}

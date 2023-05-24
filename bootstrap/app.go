package bootstrap

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Application struct {
	Env      *Env
	Database *gorm.DB
	Cors     cors.Config
}

func NewApp() *Application {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := &Application{}
	app.Env = NewEnv()
	app.Database, err = NewPostgresDatabase(app.Env)
	app.Cors = NewCorsConfig(app.Env)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connection established.")

	return app
}

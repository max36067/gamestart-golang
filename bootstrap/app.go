package bootstrap

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Application struct {
	Env      *Env
	Database *gorm.DB
}

func NewApp() *Application {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := &Application{}
	app.Env = NewEnv()
	app.Database, err = NewPostgresDatabase(app.Env)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connection established.")

	return app
}

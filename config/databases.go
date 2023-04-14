package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type DataBaseConfig struct {
	UserName     string
	UserPassword string
	Host         string
	Port         string
	Database     string
}

func (dbc *DataBaseConfig) InitConfig() {
	dbc.UserName = os.Getenv("DB_USERNAME")
	dbc.UserPassword = os.Getenv("DB_PASSWORD")
	dbc.Host = os.Getenv("DB_HOST")
	dbc.Port = os.Getenv("DB_PORT")
	dbc.Database = os.Getenv("DB_DATABASE")
}

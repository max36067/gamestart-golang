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

func InitConfig() *DataBaseConfig {
	return &DataBaseConfig{
		UserName:     os.Getenv("DB_USERNAME"),
		UserPassword: os.Getenv("DB_PASSWORD"),
		Host:         os.Getenv("DB_HOST"),
		Port:         os.Getenv("DB_PORT"),
		Database:     os.Getenv("DB_DATABASE"),
	}
}

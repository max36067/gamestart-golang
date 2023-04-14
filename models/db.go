package models

import (
	"apigee-portal/v2/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Session *gorm.DB

func init() {
	var dbConfig config.DataBaseConfig
	dbConfig.InitConfig()

	dbUri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbConfig.Host, dbConfig.UserName, dbConfig.UserPassword, dbConfig.Database, dbConfig.Port)

	session, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	Session = session
}

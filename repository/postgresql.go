package repository

import (
	"apigee-portal/v2/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Session *gorm.DB

func init() {
	var err error
	dbConfig := config.InitConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbConfig.Host, dbConfig.UserName, dbConfig.UserPassword, dbConfig.Database, dbConfig.Port)
	Session, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Database connect error: %v", err)
	}

	if Session.Error != nil {
		log.Fatalf("Database error %v", Session.Error)
	}

}

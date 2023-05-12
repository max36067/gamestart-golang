package bootstrap

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDatabase(env *Env) (*gorm.DB, error) {

	username := env.DBUser
	password := env.DBPassword
	host := env.DBHost
	port := env.DBPort
	dbname := env.DBName

	uri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbname, port)
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	return db, err
}

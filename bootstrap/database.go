package bootstrap

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDatabase(env *Env) (*gorm.DB, error) {

	username := env.DB_User
	password := env.DB_Password
	host := env.DB_Host
	port := env.DB_Port
	dbname := env.DB_Name

	uri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbname, port)
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	return db, err
}

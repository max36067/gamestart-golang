package bootstrap

import (
	"apigee-portal/v2/postgres"
	"fmt"
)

func NewPostgresDatabase(env *Env) postgres.DataBase {

	username := env.DBUser
	password := env.DBPassword
	host := env.DBHost
	port := env.DBPort
	dbname := env.DBName

	uri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbname, port)
	db := postgres.NewDatabase(uri)

	return db
}

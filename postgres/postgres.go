package postgres

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataBase struct {
	*gorm.DB
}

func NewDatabase(uri string) DataBase {

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connection established.")

	return DataBase{
		DB: db,
	}

}

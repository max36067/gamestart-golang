package repository

import (
	"apigee-portal/v2/domain"

	"gorm.io/gorm"
)

type saltRepository struct {
	database *gorm.DB
}

func NewSaltRepository(db *gorm.DB) domain.SaltRepository {
	return &saltRepository{
		database: db,
	}
}

func (sr *saltRepository) GetSaltByEmail(email string) (saltString string, err error) {
	var salt domain.Salt
	err = sr.database.First(&salt, "email = ?", email).Error

	return salt.Salt, err
}

func (sr *saltRepository) Create(salt *domain.Salt) error {
	return sr.database.Create(salt).Error
}

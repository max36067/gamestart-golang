package repository

import (
	"apigee-portal/v2/domain"
	"apigee-portal/v2/postgres"
)

type userRepository struct {
	database postgres.DataBase
}

func NewUserRepository(db postgres.DataBase) domain.UserRepository {
	return &userRepository{
		database: db,
	}
}

func (ur *userRepository) Create(user *domain.User) error {
	return ur.database.Create(&user).Error
}

func (ur *userRepository) Fetch() ([]domain.UserResponse, error) {
	var users []domain.UserResponse
	if err := ur.database.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *userRepository) GetByEmail(email string) (domain.UserResponse, error) {
	var user domain.UserResponse
	err := ur.database.Select(&user).Where("email = ?", email).Error
	return user, err
}

func (ur *userRepository) GetByID(id string) (domain.UserResponse, error) {
	var user domain.UserResponse
	err := ur.database.Select(&user).Where("id = ?", id).Error
	return user, err
}

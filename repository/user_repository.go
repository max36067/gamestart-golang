package repository

import (
	"apigee-portal/v2/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
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

func (ur *userRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	if err := ur.database.First(&user, "email = ?", email).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (ur *userRepository) GetByID(id int) (domain.User, error) {
	var user domain.User
	if err := ur.database.First(&user, "id = ?", id).Error; err != nil {
		return user, err
	}
	return user, nil
}

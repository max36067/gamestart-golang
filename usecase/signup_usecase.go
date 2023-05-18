package usecase

import (
	"apigee-portal/v2/domain"
	"apigee-portal/v2/utils"
	"time"
)

type signupUsecase struct {
	userRepository domain.UserRepository
	saltRepository domain.SaltRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(userRepository domain.UserRepository, saltRepository domain.SaltRepository, contextTimeout time.Duration) domain.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		saltRepository: saltRepository,
		contextTimeout: contextTimeout,
	}
}

func (su *signupUsecase) CreateUser(user *domain.User) error {
	return su.userRepository.Create(user)
}

func (su *signupUsecase) CreateSalt(salt *domain.Salt) error {
	return su.saltRepository.Create(salt)
}

func (su *signupUsecase) GetUserByEmail(email string) (domain.User, error) {
	return su.userRepository.GetByEmail(email)
}

func (su *signupUsecase) CreateAccessToken(user *domain.User, secret string, expire int) (accessToken string, err error) {
	return utils.GenerateAccessToken(user, secret, expire)
}

func (su *signupUsecase) CreateRefreshToken(user *domain.User, secret string, expire int) (refreshToken string, err error) {
	return utils.GenerateRefreshToken(user, secret, expire)
}

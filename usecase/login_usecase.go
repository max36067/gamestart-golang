package usecase

import (
	"apigee-portal/v2/domain"
	"apigee-portal/v2/utils"
	"time"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginUsecase) GetUserByEmail(email string) (domain.User, error) {
	return lu.userRepository.GetByEmail(email)
}

func (lu *loginUsecase) CreateAccessToken(user *domain.User, secret string, expire int) (access_token string, err error) {
	return utils.GenerateAccessToken(user, secret, expire)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expire int) (refreshToken string, err error) {
	return utils.GenerateRefreshToken(user, secret, expire)
}

func (lu *loginUsecase) VerifyPassword(hashedPassword, candidatePassword string) error {
	return utils.VerifyPassword(hashedPassword, candidatePassword)
}

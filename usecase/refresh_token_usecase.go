package usecase

import (
	"apigee-portal/v2/domain"
	"apigee-portal/v2/utils"
	"time"
)

type refreshTokenUsecase struct {
	userRepository domain.UserRepository
	timeout        time.Duration
}

func NewRefreshTokenUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepository: userRepository,
		timeout:        timeout,
	}
}

func (rtu *refreshTokenUsecase) GetUserByID(id int) (domain.User, error) {
	return rtu.userRepository.GetByID(id)
}

func (rtu *refreshTokenUsecase) CreateAccessToken(user *domain.User, secret string, expire int) (accessToken string, err error) {
	return utils.GenerateAccessToken(user, secret, expire)
}

func (rtu *refreshTokenUsecase) CreateRefreshToken(user *domain.User, secret string, expire int) (refreshToken string, err error) {
	return utils.GenerateRefreshToken(user, secret, expire)
}

func (rtu *refreshTokenUsecase) ExtractIDFromToken(requestToken string, secret string) (int, error) {
	return utils.ExtractIDFromToken(requestToken, secret)
}

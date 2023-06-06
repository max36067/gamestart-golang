package usecase

import (
	"apigee-portal/v2/domain"
	"context"
	"time"
)

type logoutUsecase struct {
	tokenBlacklistRepository domain.TokenBlacklistRepository
	timeout                  time.Duration
}

func NewLogoutUsecase(tokenBlacklistRepository domain.TokenBlacklistRepository, timeout time.Duration) domain.LogoutUsecase {
	return &logoutUsecase{
		timeout: timeout,
	}
}

func (lu *logoutUsecase) RevokeRefreshToken(ctx context.Context, jwtTokenBlacklist *domain.JWTTokenBlacklist, exp time.Duration) error {
	return lu.tokenBlacklistRepository.SetRefreshToken(ctx, jwtTokenBlacklist, exp)
}

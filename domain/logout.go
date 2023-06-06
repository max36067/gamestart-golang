package domain

import (
	"context"
	"time"
)

type LogoutUsecase interface {
	RevokeRefreshToken(ctx context.Context, jwtTokenBlacklist *JWTTokenBlacklist, exp time.Duration) error
}

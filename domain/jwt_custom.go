package domain

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
	jwt.RegisteredClaims
}

type JWTCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

type JWTTokenBlacklist struct {
	TokenKey string
	Token    string
}

type TokenBlacklistRepository interface {
	SetRefreshToken(ctx context.Context, jwtTokenBlacklist *JWTTokenBlacklist, exp time.Duration) error
}

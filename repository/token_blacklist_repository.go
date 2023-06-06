package repository

import (
	"apigee-portal/v2/domain"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type tokenBlacklistRepository struct {
	rdb *redis.Client
}

func NewTokenBlacklistRepository(rdb *redis.Client) domain.TokenBlacklistRepository {
	return &tokenBlacklistRepository{
		rdb: rdb,
	}
}

func (tbr *tokenBlacklistRepository) SetRefreshToken(ctx context.Context, jwtTokenBlacklist *domain.JWTTokenBlacklist, exp time.Duration) error {
	return tbr.rdb.Set(ctx, jwtTokenBlacklist.TokenKey, jwtTokenBlacklist.Token, exp).Err()
}

func (tbr *tokenBlacklistRepository) GetRefreshToken(ctx context.Context, jwtTokenBlacklist *domain.JWTTokenBlacklist) (string, error) {
	return tbr.rdb.Get(ctx, jwtTokenBlacklist.TokenKey).Result()

}

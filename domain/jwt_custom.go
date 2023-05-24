package domain

import (
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

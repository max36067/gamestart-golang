package utils

import (
	"apigee-portal/v2/domain"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(user *domain.User, secret string, expire int) (accessToken string, err error) {

	exp := time.Now().Add(time.Minute * time.Duration(expire))
	claims := &domain.JWTCustomClaims{
		Name: user.Name,
		ID:   user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func GenerateRefreshToken(user *domain.User, secret string, expire int) (refreshToken string, err error) {
	exp := time.Now().Add(time.Minute * time.Duration(expire))
	claims := &domain.JWTCustomClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func ExtractIDFromToken(requestToken string, secret string) (int, error) {
	token, err := jwt.ParseWithClaims(requestToken, &domain.JWTCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claim, ok := token.Claims.(*domain.JWTCustomClaims); ok && token.Valid {
		return claim.ID, nil
	} else {
		return 0, err
	}
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

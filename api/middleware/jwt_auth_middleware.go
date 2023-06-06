package middleware

import (
	"apigee-portal/v2/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func JWTAuthMiddleware(rdb *redis.Client, secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")

		if len(t) != 2 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
			ctx.Abort()
			return
		}
		authToken := t[1]
		authorized, err := utils.IsAuthorized(authToken, secret)
		if !authorized {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			ctx.Abort()
			return
		}
		userID, err := utils.ExtractIDFromToken(authToken, secret)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Set("x-user-id", userID)

		cookie, err := ctx.Request.Cookie("refresh_token")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Refresh token not found."})
			ctx.Abort()
			return
		}
		refreshToken := cookie.Value

		_, err = rdb.Get(ctx, fmt.Sprint("bl_%s", refreshToken)).Result()
		if err == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Refresh token is invalid"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

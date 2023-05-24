package middleware

import (
	"apigee-portal/v2/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")

		if len(t) == 2 {
			authToken := t[1]
			authorized, err := utils.IsAuthorized(authToken, secret)
			if authorized {
				userID, err := utils.ExtractIDFromToken(authToken, secret)
				if err != nil {
					ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
					ctx.Abort()
					return
				}
				ctx.Set("x-user-id", userID)
				ctx.Next()
				return
			}
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		ctx.Abort()
	}
}

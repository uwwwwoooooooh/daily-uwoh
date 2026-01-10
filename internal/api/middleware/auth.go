package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/token"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/utils"
)

func AuthMiddleware(tokenMaker token.TokenMaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.SendError(c, http.StatusUnauthorized, "Authorization header is required")
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			utils.SendError(c, http.StatusUnauthorized, "Invalid authorization header format")
			return
		}

		payload, err := tokenMaker.VerifyToken(bearerToken[1])
		if err != nil {
			utils.SendError(c, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		c.Set("userID", payload.UserID)
		c.Next()
	}
}

package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/config"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/utils"
)

func AuthMiddleware(cfg config.Config) gin.HandlerFunc {
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

		claims, err := utils.ValidateToken(bearerToken[1], cfg.JWTSecret)
		if err != nil {
			utils.SendError(c, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}

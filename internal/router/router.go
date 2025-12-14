package router

import (
	"github.com/gin-gonic/gin"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/handler"
)

// NewRouter initializes the Gin engine and defines routes
func NewRouter(authHandler *handler.AuthHandler) *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	return r
}

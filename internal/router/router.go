package router

import (
	"github.com/gin-gonic/gin"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/config"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/handler"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/middleware"
)

// NewRouter initializes the Gin engine and defines routes
func NewRouter(authHandler *handler.AuthHandler, cfg config.Config) *gin.Engine {
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.GET("/me", middleware.AuthMiddleware(cfg), authHandler.Me)
	}

	return r
}

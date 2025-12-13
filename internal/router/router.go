package router

import (
	"github.com/gin-gonic/gin"
)

// NewRouter initializes the Gin engine and defines routes
func NewRouter() *gin.Engine {
	r := gin.Default()

	// TODO: Define API groups
	// v1 := r.Group("/api/v1")
	// {
	// 	// v1.POST("/upload", handleUpload)
	// 	// v1.GET("/feed", handleFeed)
	// }

	return r
}

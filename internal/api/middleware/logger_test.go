package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("MiddlewareExecutes", func(t *testing.T) {
		r := gin.New()
		r.Use(Logger())

		r.GET("/test", func(c *gin.Context) {
			// Verify context variable set by logger
			example, exists := c.Get("example")
			require.True(t, exists)
			require.Equal(t, "12345", example)
			c.Status(http.StatusOK)
		})

		req, _ := http.NewRequest("GET", "/test", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
	})
}

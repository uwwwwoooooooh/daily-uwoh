package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestRecovery(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("RecoversFromPanic", func(t *testing.T) {
		r := gin.New()
		// Important: Connect output to a pipe or discard to avoid noisy logs during test
		r.Use(Recovery())

		r.GET("/panic", func(c *gin.Context) {
			panic("something went wrong")
		})

		req, _ := http.NewRequest("GET", "/panic", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusInternalServerError, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		require.Equal(t, "Internal Server Error", response["error"])
	})

	t.Run("NoPanic", func(t *testing.T) {
		r := gin.New()
		r.Use(Recovery())

		r.GET("/ok", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req, _ := http.NewRequest("GET", "/ok", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
	})
}

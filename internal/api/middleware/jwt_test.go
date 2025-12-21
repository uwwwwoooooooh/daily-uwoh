package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/config"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/utils"
)

func TestAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	cfg := config.Config{
		JWTSecret: "test_secret",
	}

	tests := []struct {
		name          string
		setupAuth     func(t *testing.T, req *http.Request)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "NoAuthorizationHeader",
			setupAuth: func(t *testing.T, req *http.Request) {
				// No header
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InvalidAuthorizationFormat",
			setupAuth: func(t *testing.T, req *http.Request) {
				req.Header.Set("Authorization", "InvalidFormat")
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InvalidToken",
			setupAuth: func(t *testing.T, req *http.Request) {
				req.Header.Set("Authorization", "Bearer invalid_token")
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "ValidToken",
			setupAuth: func(t *testing.T, req *http.Request) {
				token, err := utils.GenerateToken(1, cfg.JWTSecret, 1)
				require.NoError(t, err)
				req.Header.Set("Authorization", "Bearer "+token)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(recorder)

			// Mock request
			req, _ := http.NewRequest("GET", "/", nil)
			tc.setupAuth(t, req)
			c.Request = req

			authMiddleware := AuthMiddleware(cfg)
			handler := func(c *gin.Context) {
				c.Status(http.StatusOK)
			}

			// Manually chain setup
			// Since AuthMiddleware returns a handler, we simulate the middleware chain
			// BUT AuthMiddleware calls c.Next(), so in a unit test of just the middleware,
			// we can't easily chain without a router.
			// Easier approach: Use a router

			r := gin.New()
			r.Use(authMiddleware)
			r.GET("/", handler)
			r.ServeHTTP(recorder, req)

			tc.checkResponse(t, recorder)
		})
	}
}

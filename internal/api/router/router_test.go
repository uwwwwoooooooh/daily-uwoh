package router

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/api/handler"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/config"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/model"
)

// MockAuthService (Redefined here or could be shared if in a separate package,
// but for simplicity locally defined or we could mock the handler itself if interfaces allowed,
// but NewRouter takes *handler.AuthHandler struct likely.
// Let's re-use the mock approach or just mock the dependencies of the handler.)

// Since NewRouter takes concrete *handler.AuthHandler, we need to pass a handler
// initialized with a mock service.

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Register(ctx context.Context, email, password string) (*model.User, error) {
	args := m.Called(ctx, email, password)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockAuthService) Login(ctx context.Context, email, password string) (string, error) {
	args := m.Called(ctx, email, password)
	return args.String(0), args.Error(1)
}

func (m *MockAuthService) GetMe(ctx context.Context, userID uint) (*model.User, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func TestNewRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockAuthService)
	authHandler := handler.NewAuthHandler(mockService)
	cfg := config.Config{JWTSecret: "test"}

	r := NewRouter(authHandler, cfg)

	tests := []struct {
		method string
		path   string
	}{
		{"POST", "/auth/register"},
		{"POST", "/auth/login"},
		{"GET", "/auth/me"},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			// We optimize by just checking if the route exists in the engine's trees
			// OR by sending a request and asserting 404 is NOT returned.
			// Checking 404 is easier.

			req, _ := http.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.NotEqual(t, http.StatusNotFound, w.Code, "Route %s %s should exist", tt.method, tt.path)
		})
	}
}

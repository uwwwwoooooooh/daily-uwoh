package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/model"
)

// MockAuthService is a mock implementation of service.AuthService
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

func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		mockService := new(MockAuthService)
		user := &model.User{
			ID:        1,
			Email:     "test@example.com",
			CreatedAt: time.Now(),
		}
		mockService.On("Register", mock.Anything, "test@example.com", "password123").Return(user, nil)

		authHandler := NewAuthHandler(mockService)
		r := gin.Default()
		r.POST("/register", authHandler.Register)

		reqBody := map[string]string{
			"email":    "test@example.com",
			"password": "password123",
		}
		body, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusCreated, w.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("BadRequest", func(t *testing.T) {
		mockService := new(MockAuthService)
		authHandler := NewAuthHandler(mockService)
		r := gin.Default()
		r.POST("/register", authHandler.Register)

		reqBody := map[string]string{
			"email": "invalid-email",
		}
		body, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		mockService := new(MockAuthService)
		token := "test_token"
		mockService.On("Login", mock.Anything, "test@example.com", "password123").Return(token, nil)

		authHandler := NewAuthHandler(mockService)
		r := gin.Default()
		r.POST("/login", authHandler.Login)

		reqBody := map[string]string{
			"email":    "test@example.com",
			"password": "password123",
		}
		body, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockService := new(MockAuthService)
		mockService.On("Login", mock.Anything, "test@example.com", "wrongpassword").Return("", errors.New("invalid credentials"))

		authHandler := NewAuthHandler(mockService)
		r := gin.Default()
		r.POST("/login", authHandler.Login)

		reqBody := map[string]string{
			"email":    "test@example.com",
			"password": "wrongpassword",
		}
		body, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusUnauthorized, w.Code)
		mockService.AssertExpectations(t)
	})
}

func TestMe(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		mockService := new(MockAuthService)
		user := &model.User{
			ID:    1,
			Email: "test@example.com",
		}
		mockService.On("GetMe", mock.Anything, uint(1)).Return(user, nil)

		authHandler := NewAuthHandler(mockService)
		r := gin.Default()
		// Middleware simulation: set userID in context
		r.GET("/me", func(c *gin.Context) {
			c.Set("userID", uint(1))
			authHandler.Me(c)
		})

		req, _ := http.NewRequest(http.MethodGet, "/me", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockService := new(MockAuthService)
		authHandler := NewAuthHandler(mockService)
		r := gin.Default()
		// No middleware setting userID
		r.GET("/me", authHandler.Me)

		req, _ := http.NewRequest(http.MethodGet, "/me", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusUnauthorized, w.Code)
	})
}

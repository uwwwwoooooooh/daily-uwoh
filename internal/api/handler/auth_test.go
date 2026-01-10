package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"

	"github.com/uwwwwoooooooh/daily-uwoh/internal/api/handler"
	mockdb "github.com/uwwwwoooooooh/daily-uwoh/internal/db/mock"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/db/sqlc"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/repository"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/service"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/token"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/utils"
)

func TestLogin(t *testing.T) {
	// Setup Helper for Password
	password := "secret123"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	require.NoError(t, err)

	user := sqlc.Users{
		ID:        1,
		Email:     "test@example.com",
		Password:  string(hashedPassword),
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"email":    user.Email,
				"password": password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Eq(user.Email)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)

				// Verify JSON response structure
				var responseBody map[string]interface{}
				err := json.Unmarshal(recorder.Body.Bytes(), &responseBody)
				require.NoError(t, err)

				// Expecting {"access_token": "...", "refresh_token": "..."}
				require.NotEmpty(t, responseBody["access_token"])
				require.NotEmpty(t, responseBody["refresh_token"])
			},
		},
		{
			name: "NotFound",
			body: gin.H{
				"email":    user.Email,
				"password": password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Eq(user.Email)).
					Times(1).
					Return(sqlc.Users{}, pgx.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
				// Expecting {"error": "..."}
				var responseBody map[string]interface{}
				err := json.Unmarshal(recorder.Body.Bytes(), &responseBody)
				require.NoError(t, err)
				require.NotEmpty(t, responseBody["error"])
			},
		},
		{
			name: "InvalidPassword",
			body: gin.H{
				"email":    user.Email,
				"password": "wrong_password",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Eq(user.Email)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
				require.Contains(t, recorder.Body.String(), "error")
			},
		},
		{
			name: "BadRequest",
			body: gin.H{
				"email": "invalid-email",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			sqlStore := &repository.SQLStore{
				Store: store,
			}

			config := utils.Config{
				TokenSymmetricKey:    "12345678901234567890123456789012",
				AccessTokenDuration:  15 * time.Minute,
				RefreshTokenDuration: 7 * 24 * time.Hour,
			}

			tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
			require.NoError(t, err)

			authService := service.NewAuthService(sqlStore, tokenMaker, config)
			authHandler := handler.NewAuthHandler(authService)

			// Setup Router
			router := gin.Default()
			router.POST("/login", authHandler.Login)

			// Make Request
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(data))
			require.NoError(t, err)

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}
}

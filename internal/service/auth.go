package service

import (
	"context"

	"github.com/uwwwwoooooooh/daily-uwoh/internal/model"
)

type AuthService interface {
	Register(ctx context.Context, email, password string) (*model.User, error)
	Login(ctx context.Context, email, password string) (string, error)
}

type authService struct {
	// repo repository.UserRepository
}

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) Register(ctx context.Context, email, password string) (*model.User, error) {
	// TODO: hash password, save user
	return nil, nil
}

func (s *authService) Login(ctx context.Context, email, password string) (string, error) {
	// TODO: verify password, sign jwt
	return "token", nil
}

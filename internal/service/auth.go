package service

import (
	"context"
	"time"

	"errors"

	"github.com/uwwwwoooooooh/daily-uwoh/internal/model"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/repository"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/token"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/utils"
)

type AuthService interface {
	Register(ctx context.Context, email, password string) (*model.User, error)
	Login(ctx context.Context, email, password string) (string, error)
	GetMe(ctx context.Context, userID uint) (*model.User, error)
}

type authService struct {
	repo       repository.UserRepository
	tokenMaker token.TokenMaker
	config     utils.Config
}

func NewAuthService(repo repository.UserRepository, tokenMaker token.TokenMaker, cfg utils.Config) AuthService {
	return &authService{
		repo:       repo,
		tokenMaker: tokenMaker,
		config:     cfg,
	}
}

func (s *authService) Register(ctx context.Context, email, password string) (*model.User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:    email,
		Password: hashedPassword,
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := utils.CheckPassword(password, user.Password); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, _, err := s.tokenMaker.CreateToken(user.ID, time.Duration(s.config.AccessTokenDuration)*time.Hour)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *authService) GetMe(ctx context.Context, userID uint) (*model.User, error) {
	return s.repo.FindByID(ctx, userID)
}

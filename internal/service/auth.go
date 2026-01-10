package service

import (
	"context"

	"errors"

	"github.com/uwwwwoooooooh/daily-uwoh/internal/model"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/repository"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/token"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/utils"
)

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthService interface {
	Register(ctx context.Context, email, password string) (*model.User, error)
	Login(ctx context.Context, email, password string) (*LoginResponse, error)
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

func (s *authService) Login(ctx context.Context, email, password string) (*LoginResponse, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := utils.CheckPassword(password, user.Password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	accessToken, _, err := s.tokenMaker.CreateToken(user.ID, s.config.AccessTokenDuration)
	if err != nil {
		return nil, err
	}

	refreshToken, _, err := s.tokenMaker.CreateToken(user.ID, s.config.RefreshTokenDuration)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authService) GetMe(ctx context.Context, userID uint) (*model.User, error) {
	return s.repo.FindByID(ctx, userID)
}

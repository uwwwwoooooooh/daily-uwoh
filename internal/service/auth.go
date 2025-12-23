package service

import (
	"context"

	"errors"

	"github.com/uwwwwoooooooh/daily-uwoh/internal/model"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/repository"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, email, password string) (*model.User, error)
	Login(ctx context.Context, email, password string) (string, error)
	GetMe(ctx context.Context, userID uint) (*model.User, error)
}

type authService struct {
	repo   repository.UserRepository
	config utils.Config
}

func NewAuthService(repo repository.UserRepository, cfg utils.Config) AuthService {
	return &authService{
		repo:   repo,
		config: cfg,
	}
}

func (s *authService) Register(ctx context.Context, email, password string) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:    email,
		Password: string(hashedPassword),
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

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID, s.config.JWTSecret, s.config.JWTExpirationHours)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *authService) GetMe(ctx context.Context, userID uint) (*model.User, error) {
	return s.repo.FindByID(ctx, userID)
}

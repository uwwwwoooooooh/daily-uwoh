package repository

import (
	"context"

	"github.com/uwwwwoooooooh/daily-uwoh/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, u *model.User) error
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}

type postgresUserRepository struct {
	// db *gorm.DB
}

func NewPostgresUserRepository() UserRepository {
	return &postgresUserRepository{}
}

func (r *postgresUserRepository) Create(ctx context.Context, u *model.User) error {
	// TODO
	return nil
}

func (r *postgresUserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	// TODO
	return nil, nil
}

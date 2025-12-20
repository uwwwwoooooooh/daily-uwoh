package repository

import (
	"context"

	"github.com/uwwwwoooooooh/daily-uwoh/internal/db/sqlc"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, u *model.User) error
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	FindByID(ctx context.Context, id uint) (*model.User, error)
}

// implementation in store.go usually, but here we can define methods on SQLStore if we want,
// or keep it separate. The plan was "Refactor Repositories to use internal/db".
// Since I defined SQLStore in store.go, I should attach methods to SQLStore there or here.
// Go allows splitting methods across files in the same package.

func (s *SQLStore) CreateUser(ctx context.Context, u *model.User) error {
	params := sqlc.InsertUserParams{
		Email:    u.Email,
		Password: u.Password,
	}

	user, err := s.Queries.InsertUser(ctx, params)
	if err != nil {
		return err
	}

	u.ID = uint(user.ID)
	u.CreatedAt = user.CreatedAt.Time
	u.UpdatedAt = user.UpdatedAt.Time

	return nil
}

func (s *SQLStore) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := s.Queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:        uint(user.ID),
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	}, nil
}

func (s *SQLStore) FindByID(ctx context.Context, id uint) (*model.User, error) {
	user, err := s.Queries.GetUserByID(ctx, int64(id))
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:        uint(user.ID),
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	}, nil
}

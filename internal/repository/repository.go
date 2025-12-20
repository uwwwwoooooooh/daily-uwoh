package repository

import (
	"context"

	"github.com/uwwwwoooooooh/daily-uwoh/internal/model"
)

// ArtworkRepository defines how we interact with the database.
type ArtworkRepository interface {
	CreateArtwork(ctx context.Context, artwork *model.Artwork) error
	FindByHash(ctx context.Context, hash string) (*model.Artwork, error)
}

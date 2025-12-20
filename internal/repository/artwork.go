package repository

import (
	"context"
	"encoding/json"

	"github.com/uwwwwoooooooh/daily-uwoh/internal/db/sqlc"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/model"
)

func (s *SQLStore) CreateArtwork(ctx context.Context, artwork *model.Artwork) error {
	metaDataJSON, err := json.Marshal(artwork.MetaData)
	if err != nil {
		return err
	}

	params := sqlc.InsertArtworkParams{
		Title:     artwork.Title,
		FilePath:  artwork.FilePath,
		SourceUrl: artwork.SourceURL,
		PHash:     artwork.PHash,
		MetaData:  metaDataJSON,
		ArtistID:  int64(artwork.ArtistID),
	}

	createdArtwork, err := s.Queries.InsertArtwork(ctx, params)
	if err != nil {
		return err
	}

	artwork.ID = uint(createdArtwork.ID)
	artwork.CreatedAt = createdArtwork.CreatedAt.Time
	artwork.UpdatedAt = createdArtwork.UpdatedAt.Time

	return nil
}

func (s *SQLStore) FindByHash(ctx context.Context, hash string) (*model.Artwork, error) {
	artwork, err := s.Queries.FindByHash(ctx, hash)
	if err != nil {
		return nil, err
	}

	var metaData map[string]interface{}
	if err := json.Unmarshal(artwork.MetaData, &metaData); err != nil {
		// Log error but maybe return nil metadata or error?
		// For now return empty map if error
		metaData = make(map[string]interface{})
	}

	return &model.Artwork{
		ID:        uint(artwork.ID),
		Title:     artwork.Title,
		FilePath:  artwork.FilePath,
		SourceURL: artwork.SourceUrl,
		PHash:     artwork.PHash,
		MetaData:  metaData,
		ArtistID:  uint(artwork.ArtistID),
		CreatedAt: artwork.CreatedAt.Time,
		UpdatedAt: artwork.UpdatedAt.Time,
	}, nil
}

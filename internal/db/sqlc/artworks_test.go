package sqlc

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomArtwork(t *testing.T, artist Artists) Artworks {
	title := fmt.Sprintf("Artwork %d", time.Now().UnixNano())
	arg := InsertArtworkParams{
		Title:     title,
		FilePath:  fmt.Sprintf("/tmp/%s.jpg", title),
		SourceUrl: fmt.Sprintf("http://example.com/%s", title),
		PHash:     fmt.Sprintf("phash_%d", time.Now().UnixNano()),
		MetaData:  []byte("{}"),
		ArtistID:  artist.ID,
	}

	artwork, err := testQueries.InsertArtwork(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, artwork)

	require.Equal(t, arg.Title, artwork.Title)
	require.Equal(t, arg.ArtistID, artwork.ArtistID)
	require.Equal(t, arg.FilePath, artwork.FilePath)
	require.Equal(t, arg.SourceUrl, artwork.SourceUrl)
	require.Equal(t, arg.PHash, artwork.PHash)
	require.Equal(t, arg.MetaData, artwork.MetaData)
	require.NotZero(t, artwork.ID)

	return artwork
}

func TestCreateArtwork(t *testing.T) {
	artist := createRandomArtist(t)
	createRandomArtwork(t, artist)
}

func TestFindByHash(t *testing.T) {
	artist := createRandomArtist(t)
	artwork1 := createRandomArtwork(t, artist)

	artwork2, err := testQueries.FindByHash(context.Background(), artwork1.PHash)
	require.NoError(t, err)
	require.NotEmpty(t, artwork2)

	require.Equal(t, artwork1.ID, artwork2.ID)
	require.Equal(t, artwork1.PHash, artwork2.PHash)
	require.Equal(t, artwork1.Title, artwork2.Title)
	require.Equal(t, artwork1.ArtistID, artwork2.ArtistID)
}

func TestListArtworksByArtist(t *testing.T) {
	artist := createRandomArtist(t)
	for i := 0; i < 5; i++ {
		createRandomArtwork(t, artist)
	}

	artworks, err := testQueries.ListArtworksByArtist(context.Background(), artist.ID)
	require.NoError(t, err)
	require.Len(t, artworks, 5)

	for _, artwork := range artworks {
		require.Equal(t, artist.ID, artwork.ArtistID)
	}
}

func TestUpdateArtworkMetadata(t *testing.T) {
	artist := createRandomArtist(t)
	artwork1 := createRandomArtwork(t, artist)

	newMetaData := []byte(`{"updated": true}`)
	arg := UpdateArtworkMetadataParams{
		ID:       artwork1.ID,
		MetaData: newMetaData,
	}

	err := testQueries.UpdateArtworkMetadata(context.Background(), arg)
	require.NoError(t, err)

	artwork2, err := testQueries.FindByHash(context.Background(), artwork1.PHash)
	require.NoError(t, err)
	require.NotEmpty(t, artwork2)

	require.Equal(t, string(newMetaData), string(artwork2.MetaData))
}

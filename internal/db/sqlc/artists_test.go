package sqlc

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomArtist(t *testing.T) Artists {
	name := fmt.Sprintf("Artist_%d", time.Now().UnixNano())
	arg := InsertArtistParams{
		Name:           name,
		SocialProfiles: []byte("[]"),
	}

	artist, err := testQueries.InsertArtist(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, artist)

	require.Equal(t, arg.Name, artist.Name)
	require.Equal(t, arg.SocialProfiles, artist.SocialProfiles)
	require.NotZero(t, artist.ID)
	require.True(t, artist.CreatedAt.Valid)

	return artist
}

func TestCreateArtist(t *testing.T) {
	createRandomArtist(t)
}

func TestGetArtist(t *testing.T) {
	artist1 := createRandomArtist(t)
	artist2, err := testQueries.GetArtist(context.Background(), artist1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, artist2)

	require.Equal(t, artist1.ID, artist2.ID)
	require.Equal(t, artist1.Name, artist2.Name)
	require.Equal(t, artist1.SocialProfiles, artist2.SocialProfiles)
	require.Equal(t, artist1.CreatedAt, artist2.CreatedAt)
	require.Equal(t, artist1.UpdatedAt, artist2.UpdatedAt)
	require.Equal(t, artist1.DeletedAt, artist2.DeletedAt)
}

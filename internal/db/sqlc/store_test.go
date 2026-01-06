package sqlc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewStore(t *testing.T) {
	store := NewStore(testDB)
	require.NotEmpty(t, store)
}

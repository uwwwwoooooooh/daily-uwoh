-- name: InsertArtwork :one
INSERT INTO artworks (
  title, file_path, source_url, p_hash, meta_data, artist_id
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: ListArtworksByArtist :many
SELECT * FROM artworks
WHERE artist_id = $1
ORDER BY created_at DESC;

-- name: UpdateArtworkMetadata :exec
UPDATE artworks
SET meta_data = $2,
    updated_at = NOW()
WHERE id = $1;

-- name: FindByHash :one
SELECT * FROM artworks
WHERE p_hash = $1 LIMIT 1;

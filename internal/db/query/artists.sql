-- name: InsertArtist :one
INSERT INTO artists (
  name, social_profiles
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetArtist :one
SELECT * FROM artists
WHERE id = $1 LIMIT 1;

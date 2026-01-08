package model

import "time"

type Artwork struct {
	ID        uint                   `json:"id"`
	Title     string                 `json:"title"`
	FilePath  string                 `json:"file_path"`
	SourceURL string                 `json:"source_url"`
	PHash     string                 `json:"p_hash"`
	MetaData  map[string]interface{} `json:"meta_data"`
	ArtistID  uint                   `json:"artist_id"`
	Tags      []*Tag                 `json:"tags,omitempty"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
	DeletedAt *time.Time             `json:"deleted_at,omitempty"`
}

// Tag for categorization.
type Tag struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Category  string     `json:"category"`
	Artworks  []*Artwork `json:"artworks,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

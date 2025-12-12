package model

import (
	"gorm.io/gorm"
)

// Artwork represents the core entity of the system.
type Artwork struct {
	gorm.Model
	Title     string                 `json:"title"`
	FilePath  string                 `json:"file_path"`
	SourceURL string                 `json:"source_url"`
	PHash     string                 `json:"p_hash" gorm:"index"`
	MetaData  map[string]interface{} `gorm:"type:jsonb;serializer:json" json:"meta_data"`
	ArtistID  uint                   `json:"artist_id"`
	Tags      []*Tag                 `gorm:"many2many:artwork_tags;" json:"tags,omitempty"`
}

// Tag for categorization.
// TODO: Determine if we need a 'Confidence' field in the join table for AI confidence levels.
type Tag struct {
	gorm.Model
	Name     string     `gorm:"uniqueIndex" json:"name"`
	Category string     `json:"category"`
	Artworks []*Artwork `gorm:"many2many:artwork_tags;" json:"artworks,omitempty"`
}

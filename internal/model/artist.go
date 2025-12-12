package model

import (
	"gorm.io/gorm"
)

// Artist represents a creator on platforms like Twitter/Pixiv.
type Artist struct {
	gorm.Model
	Name           string                 `json:"name"`
	SocialProfiles map[string]interface{} `gorm:"type:jsonb;serializer:json" json:"social_profiles"`
	Artworks       []Artwork              `json:"artworks,omitempty"`
}

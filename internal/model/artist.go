package model

import "time"

type Artist struct {
	ID             uint                   `json:"id"`
	Name           string                 `json:"name"`
	SocialProfiles map[string]interface{} `json:"social_profiles"`
	Artworks       []Artwork              `json:"artworks,omitempty"`
	CreatedAt      time.Time              `json:"created_at"`
	UpdatedAt      time.Time              `json:"updated_at"`
	DeletedAt      *time.Time             `json:"deleted_at,omitempty"`
}

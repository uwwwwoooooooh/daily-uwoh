package service

import "context"

// AIProcessor defines the interface for image analysis.
type AIProcessor interface {
	// IsAnime checks if the image is an anime style illustration.
	IsAnime(ctx context.Context, imageURL string) (bool, error)

	// AssessQuality returns a "Uwoh" score (0.0 - 1.0) and descriptive tags.
	AssessQuality(ctx context.Context, imageURL string) (score float64, tags []string, err error)

	// CheckNSFW returns true if the image contains explicit content.
	CheckNSFW(ctx context.Context, imageURL string) (bool, error)
}

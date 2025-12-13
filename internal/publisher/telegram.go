package publisher

import "context"

// Publisher defines the interface for sending content to external platforms.
type Publisher interface {
	// SendImage sends an image to the configured channel.
	SendImage(ctx context.Context, imageURL string, caption string) error

	// SendMessage sends a text message.
	SendMessage(ctx context.Context, text string) error
}

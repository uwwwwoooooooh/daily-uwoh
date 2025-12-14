package publisher

import "context"

// TelegramPublisher implements the service.Publisher interface.
type TelegramPublisher struct {
	BotToken string
	ChatID   string
}

// NewTelegramPublisher creates a new instance of TelegramPublisher.
func NewTelegramPublisher(token, chatID string) *TelegramPublisher {
	return &TelegramPublisher{
		BotToken: token,
		ChatID:   chatID,
	}
}

// SendImage sends an image to the configured channel.
func (t *TelegramPublisher) SendImage(ctx context.Context, imageURL string, caption string) error {
	// TODO: Implement actual Telegram API call
	return nil
}

// SendMessage sends a text message.
func (t *TelegramPublisher) SendMessage(ctx context.Context, text string) error {
	// TODO: Implement actual Telegram API call
	return nil
}

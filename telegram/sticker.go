package telegram

import "errors"

// Sticker represents a Telegram sticker.
type Sticker struct{}

// Get fetches a Telegram sticker by ID.
func (s *Sticker) Get(id string) (string, error) {
	if id == "" {
		return "", errors.New("Invalid sticker ID")
	}
	return "", nil
}

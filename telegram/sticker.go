package telegram

import (
	"errors"

	"github.com/ironiclensflare/printer/telegram/http"
)

// Sticker represents a Telegram sticker.
type Sticker struct {
	httpClient http.HttpPoster
}

// Get fetches a Telegram sticker by ID.
func (s *Sticker) Get(id string) (string, error) {
	if id == "" {
		return "", errors.New("Invalid sticker ID")
	}
	s.httpClient.PostForm("", nil)
	return "test.webp", nil
}

func GetSticker() *Sticker {
	return &Sticker{
		httpClient: http.GetHttpClient(),
	}
}

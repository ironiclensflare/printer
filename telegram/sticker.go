package telegram

import (
	"errors"
	"net/url"
	"os"

	"github.com/ironiclensflare/printer/telegram/http"
)

// Sticker represents a Telegram sticker.
type Sticker struct {
	httpClient http.HttpPoster
}

// Get fetches a Telegram sticker by ID.
func (s *Sticker) Get(stickerID string) (string, error) {
	if stickerID == "" {
		return "", errors.New("Invalid sticker ID")
	}

	fileID := s.getFileID(stickerID)
	stickerPath := s.downloadSticker(fileID)
	return stickerPath, nil
}

func (s *Sticker) getFileID(stickerID string) string {
	const endpoint string = ""
	getFileEndpoint := endpoint + getBotToken()
	values := url.Values{}
	values.Add("file_id", stickerID)
	s.httpClient.PostForm(getFileEndpoint, values)
	return stickerID
}

func (s *Sticker) downloadSticker(fileID string) string {
	const endpoint string = "https://api.telegram.org/file/bot"
	stickerURL := endpoint + getBotToken() + "/" + fileID
	s.httpClient.Get(stickerURL)
	return fileID + ".webp"
}

func getBotToken() string {
	return os.Getenv("TELEGRAM_BOT_TOKEN")
}

// GetSticker returns an instance of Sticker.
func GetSticker() *Sticker {
	return &Sticker{
		httpClient: http.GetHttpClient(),
	}
}

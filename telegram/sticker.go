package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
	fmt.Println("Entering Sticker.Get")
	if stickerID == "" {
		return "", errors.New("Invalid sticker ID")
	}

	fileID := s.getFileID(stickerID)
	stickerPath := s.downloadSticker(fileID, stickerID)
	return stickerPath, nil
}

func (s *Sticker) getFileID(stickerID string) string {
	const endpoint string = "https://api.telegram.org/bot"
	getFileEndpoint := endpoint + getBotToken() + "/getFile"
	values := url.Values{}
	values.Add("file_id", stickerID)
	resp, _ := s.httpClient.PostForm(getFileEndpoint, values)
	body, _ := ioutil.ReadAll(resp.Body)

	type StickerDetails struct {
		File_ID        string
		File_Unique_ID string
		File_Size      int
		File_Path      string
	}
	type StickerResponse struct {
		OK     bool
		Result StickerDetails
	}

	var stickerResponse StickerResponse
	err := json.Unmarshal(body, &stickerResponse)
	if err != nil {
		fmt.Println("Error unmarshalling JSON", err)
	}
	fmt.Println(stickerResponse)
	return stickerResponse.Result.File_Path
}

func (s *Sticker) downloadSticker(fileID string, stickerID string) string {
	const endpoint string = "https://api.telegram.org/file/bot"
	stickerURL := endpoint + getBotToken() + "/" + fileID
	fmt.Println("Attempting to download sticker at " + stickerURL)
	stickerFileName := stickerID + ".png"
	fileName, _ := s.httpClient.DownloadFile(stickerURL, stickerFileName)
	return fileName
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

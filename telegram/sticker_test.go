package telegram

import (
	"errors"
	"testing"

	"github.com/ironiclensflare/printer/telegram/fakes"
	"github.com/stretchr/testify/assert"
)

func TestGetStickerNoId(t *testing.T) {
	sticker := Sticker{}
	_, err := sticker.Get("")
	expectedError := errors.New("Invalid sticker ID")

	if assert.Error(t, err) {
		assert.Equal(t, expectedError, err)
	}
}

func TestGetStickerValidId(t *testing.T) {
	sticker, counters := getTestSticker()
	filename, err := sticker.Get("12345")

	assert.NoError(t, err)
	assert.Equal(t, "12345.webp", filename)
	assert.Equal(t, 1, *counters.GetCounter)
	assert.Equal(t, 1, *counters.PostFormCounter)
}

func getTestSticker() (*Sticker, *Counters) {
	fakeHTTPClient := fakes.GetFakeHttpClient()
	sticker := Sticker{httpClient: fakeHTTPClient}
	counters := Counters{}
	counters.GetCounter = &fakeHTTPClient.GetCounter
	counters.PostFormCounter = &fakeHTTPClient.PostFormCounter
	return &sticker, &counters
}

type Counters struct {
	GetCounter      *int
	PostFormCounter *int
}

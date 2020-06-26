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
	sticker, counter := getTestSticker()
	filename, err := sticker.Get("12345")

	assert.NoError(t, err)
	assert.Equal(t, "test.webp", filename)
	assert.Equal(t, 1, *counter)
}

func getTestSticker() (*Sticker, *int) {
	fakeHttpClient := fakes.GetFakeHttpClient()
	sticker := Sticker{httpClient: fakeHttpClient}
	return &sticker, &fakeHttpClient.PostFormCounter
}

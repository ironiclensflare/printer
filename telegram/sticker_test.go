package telegram

import (
	"errors"
	"testing"

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
	t.Skip("Incomplete test")

	sticker := Sticker{}
	filename, error := sticker.Get("12345")

	assert.Nil(t, error)
	assert.Equal(t, "test.webp", filename)
}

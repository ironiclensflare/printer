package telegram

import (
	"testing"
)

func TestGetStickerNoId(t *testing.T) {
	sticker := Sticker{}
	_, err := sticker.Get("")
	if err != nil {
		return
	}
	t.Error("Did not throw error")
}

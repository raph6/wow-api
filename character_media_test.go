package wowapi_test

import (
	"testing"

	"github.com/raph6/wowapi"
)

func TestCharacterCharacterMedia(t *testing.T) {
	// Test with a valid character
	var character_mediaStruct wowapi.CharacterMedia
	character_media, err := req.CharacterCharacterMedia("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character character_media: %v", err)
	}
	if character_media.Character == character_mediaStruct.Character {
		t.Error("No character_media returned")
	}

	// Test with an invalid character
	character_media, err = req.CharacterCharacterMedia("invalidrealm", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if character_media.Character != character_mediaStruct.Character {
		t.Error("Expected no character_media, but some were returned")
	}
}

package wowapi_test

import (
	"testing"

	"github.com/raph6/wowapi"
)

func TestCharacterInfo(t *testing.T) {
	// Test with a valid character
	var characterStruct wowapi.Character
	character, err := req.CharacterInfo("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character character: %v", err)
	}
	if character == characterStruct {
		t.Error("No character returned")
	}

	// Test with an invalid character
	character, err = req.CharacterInfo("invalidrealm", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if character != characterStruct {
		t.Error("Expected no character, but some were returned")
	}
}

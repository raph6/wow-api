package wowapi_test

import (
	"testing"

	"github.com/raph6/wowapi"
)

func TestCharacterAppearance(t *testing.T) {
	// Test with a valid character
	var appearanceStruct wowapi.Appearance
	appearance, err := req.CharacterAppearance("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character appearance: %v", err)
	}
	if appearance.PlayableClass == appearanceStruct.PlayableClass {
		t.Error("No appearance returned")
	}

	// Test with an invalid character
	appearance, err = req.CharacterAppearance("invalidrealm", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if appearance.PlayableClass != appearanceStruct.PlayableClass {
		t.Error("Expected no appearance, but some were returned")
	}
}

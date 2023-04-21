package wowapi_test

import (
	"testing"

	"github.com/raph6/wowapi"
)

func TestCharacterTitles(t *testing.T) {
	// Test with a valid character
	var titlesStruct wowapi.Titles
	titles, err := req.CharacterTitles("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character titles: %v", err)
	}
	if titles.Links == titlesStruct.Links {
		t.Error("No titles returned")
	}

	// Test with an invalid character
	titles, err = req.CharacterTitles("invalidrealm", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if titles.Links != titlesStruct.Links {
		t.Error("Expected no titles, but some were returned")
	}
}

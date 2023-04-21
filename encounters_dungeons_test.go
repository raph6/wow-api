package wowapi_test

import (
	"testing"

	"github.com/raph6/wowapi"
)

func TestCharacterEncountersDungeons(t *testing.T) {
	// Test with a valid character
	var encounters_dungeonsStruct wowapi.EncountersDungeons
	encounters_dungeons, err := req.CharacterEncountersDungeons("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character encounters_dungeons: %v", err)
	}
	if encounters_dungeons.Links == encounters_dungeonsStruct.Links {
		t.Error("No encounters_dungeons returned")
	}

	// Test with an invalid character
	encounters_dungeons, err = req.CharacterEncountersDungeons("invalidrealm", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if encounters_dungeons.Links != encounters_dungeonsStruct.Links {
		t.Error("Expected no encounters_dungeons, but some were returned")
	}
}

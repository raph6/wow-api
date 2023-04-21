package wowapi_test

import (
	"testing"

	"github.com/raph6/wowapi"
)

func TestCharacterQuests(t *testing.T) {
	// Test with a valid character
	var questsStruct wowapi.Quests
	quests, err := req.CharacterQuests("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character quests: %v", err)
	}
	if quests.Links == questsStruct.Links {
		t.Error("No quests returned")
	}

	// Test with an invalid character
	quests, err = req.CharacterQuests("invalidrealm", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if quests.Links != questsStruct.Links {
		t.Error("Expected no quests, but some were returned")
	}
}

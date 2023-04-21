package wowapi_test

import (
	"testing"

	"github.com/raph6/wowapi"
)

func TestCharacterQuestsCompleted(t *testing.T) {
	// Test with a valid character
	var quests_completedStruct wowapi.QuestsCompleted
	quests_completed, err := req.CharacterQuestsCompleted("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character quests_completed: %v", err)
	}
	if quests_completed.Character == quests_completedStruct.Character {
		t.Error("No quests_completed returned")
	}

	// Test with an invalid character
	quests_completed, err = req.CharacterQuestsCompleted("invalidrealm", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if quests_completed.Character != quests_completedStruct.Character {
		t.Error("Expected no quests_completed, but some were returned")
	}
}

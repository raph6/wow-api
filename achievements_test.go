package wowapi_test

import (
	"testing"

	"github.com/raph6/wowapi"
)

func TestCharacterAchievements(t *testing.T) {
	// Test with a valid character
	var achievementsStruct wowapi.Achievements
	achievements, err := req.CharacterAchievements("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character achievements: %v", err)
	}
	if achievements.TotalQuantity == achievementsStruct.TotalQuantity {
		t.Error("No achievements returned")
	}

	// Test with an invalid character
	achievements, err = req.CharacterAchievements("invalidrealm", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if achievements.TotalQuantity != achievementsStruct.TotalQuantity {
		t.Error("Expected no achievements, but some were returned")
	}
}

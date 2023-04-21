package wowapi_test

import (
	"testing"

	"github.com/raph6/wowapi"
)

func TestCharacterAchievementsStatistics(t *testing.T) {
	// Test with a valid character
	var achievements_statisticsStruct wowapi.AchievementsStatistics
	achievements_statistics, err := req.CharacterAchievementsStatistics("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character achievements_statistics: %v", err)
	}
	if achievements_statistics.Character == achievements_statisticsStruct.Character {
		t.Error("No achievements_statistics returned")
	}

	// Test with an invalid character
	achievements_statistics, err = req.CharacterAchievementsStatistics("invalidrealm", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if achievements_statistics.Character != achievements_statisticsStruct.Character {
		t.Error("Expected no achievements_statistics, but some were returned")
	}
}

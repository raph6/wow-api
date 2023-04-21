package wowapi_test

import (
	"testing"

	"github.com/raph6/wowapi"
)

func TestCharacterStatistics(t *testing.T) {
	// Test with a valid character
	var statsStruct wowapi.Statistics
	stats, err := req.CharacterStatistics("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character stats: %v", err)
	}
	if stats == statsStruct {
		t.Error("No stats returned")
	}

	// Test with an invalid character
	stats, err = req.CharacterStatistics("invalidrealm", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if stats != statsStruct {
		t.Error("Expected no stats, but some were returned")
	}
}

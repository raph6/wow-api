package wowapi_test

import (
	"testing"

	"github.com/raph6/wowapi"
)

func TestCharacterStatistics(t *testing.T) {
	// Test with a valid character
	var statisticsStruct wowapi.Statistics
	statistics, err := req.CharacterStatistics("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character statistics: %v", err)
	}
	if statistics == statisticsStruct {
		t.Error("No statistics returned")
	}

	// Test with an invalid character
	statistics, err = req.CharacterStatistics("invalidrealm", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if statistics != statisticsStruct {
		t.Error("Expected no statistics, but some were returned")
	}
}

package wowapi_test

import (
	"testing"
)

func TestCharacterStats(t *testing.T) {
	// Test with a valid character
	stats, err := req.CharacterStats("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character stats: %v", err)
	}
	if stats == nil {
		t.Error("No stats returned")
	}

	// Test with an invalid character
	stats, err = req.CharacterStats("invalidrealmname", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if stats != nil {
		t.Error("Expected no stats, but some were returned")
	}
}

package wowapi_test

import (
	"testing"

	"github.com/raph6/wowapi"
)

func TestCharacterProfessions(t *testing.T) {
	// Test with a valid character
	var professionsStruct wowapi.Professions
	professions, err := req.CharacterProfessions("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character professions: %v", err)
	}
	if professions.Links == professionsStruct.Links {
		t.Error("No professions returned")
	}

	// Test with an invalid character
	professions, err = req.CharacterProfessions("invalidrealm", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if professions.Links != professionsStruct.Links {
		t.Error("Expected no professions, but some were returned")
	}
}

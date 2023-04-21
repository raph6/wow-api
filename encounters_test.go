package wowapi_test

import (
	"testing"

	"github.com/raph6/wowapi"
)

func TestCharacterEncounters(t *testing.T) {
	// Test with a valid character
	var encountersStruct wowapi.Encounters
	encounters, err := req.CharacterEncounters("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character encounters: %v", err)
	}
	if encounters == encountersStruct {
		t.Error("No encounters returned")
	}

	// Test with an invalid character
	encounters, err = req.CharacterEncounters("invalidrealm", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if encounters != encountersStruct {
		t.Error("Expected no encounters, but some were returned")
	}
}

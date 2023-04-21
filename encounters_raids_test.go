package wowapi_test

import (
	"testing"

	"github.com/raph6/wowapi"
)

func TestCharacterEncountersRaids(t *testing.T) {
	// Test with a valid character
	var encounters_raidsStruct wowapi.EncountersRaids
	encounters_raids, err := req.CharacterEncountersRaids("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character encounters_raids: %v", err)
	}
	if encounters_raids.Links == encounters_raidsStruct.Links {
		t.Error("No encounters_raids returned")
	}

	// Test with an invalid character
	encounters_raids, err = req.CharacterEncountersRaids("invalidrealm", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if encounters_raids.Links != encounters_raidsStruct.Links {
		t.Error("Expected no encounters_raids, but some were returned")
	}
}

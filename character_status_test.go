package wowapi_test

import (
	"testing"

	"github.com/raph6/wowapi"
)

func TestCharacterCharacterStatus(t *testing.T) {
	// Test with a valid character
	var character_statusStruct wowapi.CharacterStatus
	character_status, err := req.CharacterCharacterStatus("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character character_status: %v", err)
	}
	if character_status == character_statusStruct {
		t.Error("No character_status returned")
	}

	// Test with an invalid character
	character_status, err = req.CharacterCharacterStatus("invalidrealm", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if character_status != character_statusStruct {
		t.Error("Expected no character_status, but some were returned")
	}
}

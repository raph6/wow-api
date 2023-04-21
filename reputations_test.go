package wowapi_test

import (
	"testing"

	"github.com/raph6/wowapi"
)

func TestCharacterReputations(t *testing.T) {
	// Test with a valid character
	var reputationsStruct wowapi.Reputations
	reputations, err := req.CharacterReputations("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character reputations: %v", err)
	}
	if reputations.Links == reputationsStruct.Links {
		t.Error("No reputations returned")
	}

	// Test with an invalid character
	reputations, err = req.CharacterReputations("invalidrealm", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if reputations.Links != reputationsStruct.Links {
		t.Error("Expected no reputations, but some were returned")
	}
}

package wowapi_test

import (
	"testing"

	"github.com/raph6/wowapi"
)

func TestCharacterSpecializations(t *testing.T) {
	// Test with a valid character
	var specializationsStruct wowapi.Specializations
	specializations, err := req.CharacterSpecializations("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character specializations: %v", err)
	}
	if specializations.Links == specializationsStruct.Links {
		t.Error("No specializations returned")
	}

	// Test with an invalid character
	specializations, err = req.CharacterSpecializations("invalidrealm", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if specializations.Links != specializationsStruct.Links {
		t.Error("Expected no specializations, but some were returned")
	}
}

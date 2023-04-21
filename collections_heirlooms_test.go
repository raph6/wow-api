package wowapi_test

import (
	"testing"

	"github.com/raph6/wowapi"
)

func TestCharacterCollectionsHeirlooms(t *testing.T) {
	// Test with a valid character
	var collections_heirloomsStruct wowapi.CollectionsHeirlooms
	collections_heirlooms, err := req.CharacterCollectionsHeirlooms("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character collections_heirlooms: %v", err)
	}
	if collections_heirlooms.Links == collections_heirloomsStruct.Links {
		t.Error("No collections_heirlooms returned")
	}

	// Test with an invalid character
	collections_heirlooms, err = req.CharacterCollectionsHeirlooms("invalidrealm", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if collections_heirlooms.Links != collections_heirloomsStruct.Links {
		t.Error("Expected no collections_heirlooms, but some were returned")
	}
}

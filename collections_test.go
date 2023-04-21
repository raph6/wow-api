package wowapi_test

import (
	"testing"

	"github.com/raph6/wowapi"
)

func TestCharacterCollections(t *testing.T) {
	// Test with a valid character
	var collectionsStruct wowapi.Collections
	collections, err := req.CharacterCollections("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character collections: %v", err)
	}
	if collections == collectionsStruct {
		t.Error("No collections returned")
	}

	// Test with an invalid character
	collections, err = req.CharacterCollections("invalidrealm", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if collections != collectionsStruct {
		t.Error("Expected no collections, but some were returned")
	}
}

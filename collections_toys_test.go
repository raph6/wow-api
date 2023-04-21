package wowapi_test

import (
	"testing"

	"github.com/raph6/wowapi"
)

func TestCharacterCollectionsToys(t *testing.T) {
	// Test with a valid character
	var collections_toysStruct wowapi.CollectionsToys
	collections_toys, err := req.CharacterCollectionsToys("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character collections_toys: %v", err)
	}
	if collections_toys.Links == collections_toysStruct.Links {
		t.Error("No collections_toys returned")
	}

	// Test with an invalid character
	collections_toys, err = req.CharacterCollectionsToys("invalidrealm", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if collections_toys.Links != collections_toysStruct.Links {
		t.Error("Expected no collections_toys, but some were returned")
	}
}

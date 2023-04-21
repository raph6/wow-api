package wowapi_test

import (
	"testing"

	"github.com/raph6/wowapi"
)

func TestCharacterCollectionsMounts(t *testing.T) {
	// Test with a valid character
	var collections_mountsStruct wowapi.CollectionsMounts
	collections_mounts, err := req.CharacterCollectionsMounts("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character collections_mounts: %v", err)
	}
	if collections_mounts.Links == collections_mountsStruct.Links {
		t.Error("No collections_mounts returned")
	}

	// Test with an invalid character
	collections_mounts, err = req.CharacterCollectionsMounts("invalidrealm", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if collections_mounts.Links != collections_mountsStruct.Links {
		t.Error("Expected no collections_mounts, but some were returned")
	}
}

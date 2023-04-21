package wowapi_test

import (
	"testing"

	"github.com/raph6/wowapi"
)

func TestCharacterCollectionsPets(t *testing.T) {
	// Test with a valid character
	var collections_petsStruct wowapi.CollectionsPets
	collections_pets, err := req.CharacterCollectionsPets("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character collections_pets: %v", err)
	}
	if collections_pets.Links == collections_petsStruct.Links {
		t.Error("No collections_pets returned")
	}

	// Test with an invalid character
	collections_pets, err = req.CharacterCollectionsPets("invalidrealm", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if collections_pets.Links != collections_petsStruct.Links {
		t.Error("Expected no collections_pets, but some were returned")
	}
}

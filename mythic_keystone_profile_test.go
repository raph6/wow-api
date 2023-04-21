package wowapi_test

import (
	"testing"

	"github.com/raph6/wowapi"
)

func TestCharacterMythicKeystoneProfile(t *testing.T) {
	// Test with a valid character
	var mythic_keystone_profileStruct wowapi.MythicKeystoneProfile
	mythic_keystone_profile, err := req.CharacterMythicKeystoneProfile("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character mythic_keystone_profile: %v", err)
	}
	if mythic_keystone_profile.CurrentMythicRating == mythic_keystone_profileStruct.CurrentMythicRating {
		t.Error("No mythic_keystone_profile returned")
	}

	// Test with an invalid character
	mythic_keystone_profile, err = req.CharacterMythicKeystoneProfile("invalidrealm", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if mythic_keystone_profile.CurrentMythicRating != mythic_keystone_profileStruct.CurrentMythicRating {
		t.Error("Expected no mythic_keystone_profile, but some were returned")
	}
}

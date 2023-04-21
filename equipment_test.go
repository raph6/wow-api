package wowapi_test

import (
	"testing"

	"github.com/raph6/wowapi"
)

func TestCharacterEquipment(t *testing.T) {
	// Test with a valid character
	var equipmentStruct wowapi.Equipment
	equipment, err := req.CharacterEquipment("kirin-tor", "vimdiesel")
	if err != nil {
		t.Errorf("Failed to get character equipment: %v", err)
	}
	if equipment.Links == equipmentStruct.Links {
		t.Error("No equipment returned")
	}

	// Test with an invalid character
	equipment, err = req.CharacterEquipment("invalidrealm", "invalidcharactername")
	if err == nil {
		t.Error("Expected an error, but none was returned")
	}
	if equipment.Links != equipmentStruct.Links {
		t.Error("Expected no equipment, but some were returned")
	}
}

package wowapi

import (
	"encoding/json"
	"fmt"
)

type CollectionsPets struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Pets []struct {
		CreatureDisplay struct {
			Id  float64 `json:"id"`
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
		} `json:"creature_display"`
		Id      float64 `json:"id"`
		Species struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"species"`
		Level   float64 `json:"level"`
		Quality struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"quality"`
		Stats struct {
			BreedId float64 `json:"breed_id"`
			Health  float64 `json:"health"`
			Power   float64 `json:"power"`
			Speed   float64 `json:"speed"`
		} `json:"stats"`
	} `json:"pets"`
	UnlockedBattlePetSlots float64 `json:"unlocked_battle_pet_slots"`
}

func (req RequestFunc) CharacterCollectionsPets(realm string, name string) (s CollectionsPets, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/collections/pets", realm, name)
	body, err := req(url)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &s)
	if err != nil {
		return
	}
	return
}

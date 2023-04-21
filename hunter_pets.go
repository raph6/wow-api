package wowapi

import (
	"encoding/json"
	"fmt"
)

type HunterPets struct {
	HunterPets []struct {
		CreatureDisplay struct {
			Id  float64 `json:"id"`
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
		} `json:"creature_display"`
		Name     string  `json:"name"`
		Level    float64 `json:"level"`
		Creature struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"creature"`
		Slot float64 `json:"slot"`
	} `json:"hunter_pets"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Character struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name  string  `json:"name"`
		Id    float64 `json:"id"`
		Realm struct {
			Slug string `json:"slug"`
			Key  struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"realm"`
	} `json:"character"`
}

func (req RequestFunc) CharacterHunterPets(realm string, name string) (s HunterPets, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/hunter-pets", realm, name)
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

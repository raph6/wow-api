package wowapi

import (
	"encoding/json"
	"fmt"
)

type Soulbinds struct {
	Soulbinds []struct {
		Traits []struct {
			Trait struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string  `json:"name"`
				Id   float64 `json:"id"`
			} `json:"trait"`
			Tier         float64 `json:"tier"`
			DisplayOrder float64 `json:"display_order"`
		} `json:"traits"`
		Soulbind struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"soulbind"`
	} `json:"soulbinds"`
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
	ChosenCovenant struct {
		Id  float64 `json:"id"`
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
	} `json:"chosen_covenant"`
	RenownLevel float64 `json:"renown_level"`
}

func (req RequestFunc) CharacterSoulbinds(realm string, name string) (s Soulbinds, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/soulbinds", realm, name)
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

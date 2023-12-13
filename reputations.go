package wowapi

import (
	"encoding/json"
	"fmt"
)

type Reputations struct {
	Character struct {
		Id    float64 `json:"id"`
		Realm struct {
			Slug string `json:"slug"`
			Key  struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"realm"`
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
	} `json:"character"`
	Reputations []struct {
		Faction struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"faction"`
		Standing struct {
			Raw         float64 `json:"raw"`
			Value       float64 `json:"value"`
			Max         float64 `json:"max"`
			Tier        float64 `json:"tier"`
			Name        string  `json:"name"`
			RenownLevel int64   `json:"renown_level"`
		} `json:"standing"`
	} `json:"reputations"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

func (req RequestFunc) CharacterReputations(realm string, name string) (s Reputations, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/reputations", realm, name)
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

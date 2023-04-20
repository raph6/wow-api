package wowapi

import (
	"encoding/json"
	"fmt"
)

type Encounters struct {
	Character struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name  string  `json:"name"`
		Id    float64 `json:"id"`
		Realm struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
			Slug string  `json:"slug"`
		} `json:"realm"`
	} `json:"character"`
	Dungeons struct {
		Href string `json:"href"`
	} `json:"dungeons"`
	Raids struct {
		Href string `json:"href"`
	} `json:"raids"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

func (req RequestFunc) CharacterEncounters(realm string, name string) (s Encounters, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/encounters", realm, name)
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

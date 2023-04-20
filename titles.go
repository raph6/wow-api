package wowapi

import (
	"encoding/json"
	"fmt"
)

type Titles struct {
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
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
			Slug string  `json:"slug"`
		} `json:"realm"`
	} `json:"character"`
	ActiveTitle struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name          string  `json:"name"`
		Id            float64 `json:"id"`
		DisplayString string  `json:"display_string"`
	} `json:"active_title"`
	Titles []struct {
		Id  float64 `json:"id"`
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
	} `json:"titles"`
}

func (req RequestFunc) CharacterTitles(realm string, name string) (s Titles, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/titles", realm, name)
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

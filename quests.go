package wowapi

import (
	"encoding/json"
	"fmt"
)

type Quests struct {
	InProgress []struct {
		Name string  `json:"name"`
		Id   float64 `json:"id"`
		Key  struct {
			Href string `json:"href"`
		} `json:"key"`
	} `json:"in_progress"`
	Completed struct {
		Href string `json:"href"`
	} `json:"completed"`
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
			Id   float64 `json:"id"`
			Slug string  `json:"slug"`
			Key  struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
		} `json:"realm"`
	} `json:"character"`
}

func (req RequestFunc) CharacterQuests(realm string, name string) (s Quests, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/quests", realm, name)
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

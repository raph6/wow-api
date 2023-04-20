package wowapi

import (
	"encoding/json"
	"fmt"
)

type QuestsCompleted struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Character struct {
		Id    float64 `json:"id"`
		Realm struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
			Slug string  `json:"slug"`
		} `json:"realm"`
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
	} `json:"character"`
	Quests []struct {
		Id  float64 `json:"id"`
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
	} `json:"quests"`
}

func (req RequestFunc) CharacterQuestsCompleted(realm string, name string) (s QuestsCompleted, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/quests/completed", realm, name)
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

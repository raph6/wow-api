package wowapi

import (
	"encoding/json"
	"fmt"
)

type CharacterMedia struct {
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
	Assets []struct {
		Value string `json:"value"`
		Key   string `json:"key"`
	} `json:"assets"`
}

func (req RequestFunc) CharacterCharacterMedia(realm string, name string) (s CharacterMedia, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/character-media", realm, name)
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

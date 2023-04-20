package wowapi

import (
	"encoding/json"
	"fmt"
)

type CharacterStatus struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Id      float64 `json:"id"`
	IsValid bool    `json:"is_valid"`
}

func (req RequestFunc) CharacterCharacterStatus(realm string, name string) (s CharacterStatus, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/status", realm, name)
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

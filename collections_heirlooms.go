package wowapi

import (
	"encoding/json"
	"fmt"
)

type CollectionsHeirlooms struct {
	Heirlooms []struct {
		Heirloom struct {
			Id  float64 `json:"id"`
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
		} `json:"heirloom"`
		Upgrade struct {
			Level float64 `json:"level"`
		} `json:"upgrade"`
	} `json:"heirlooms"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

func (req RequestFunc) CharacterCollectionsHeirlooms(realm string, name string) (s CollectionsHeirlooms, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/collections/heirlooms", realm, name)
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

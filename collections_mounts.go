package wowapi

import (
	"encoding/json"
	"fmt"
)

type CollectionsMounts struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Mounts []struct {
		Mount struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"mount"`
		IsUseable bool `json:"is_useable"`
	} `json:"mounts"`
}

func (req RequestFunc) CharacterCollectionsMounts(realm string, name string) (s CollectionsMounts, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/collections/mounts", realm, name)
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

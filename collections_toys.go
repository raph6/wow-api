package wowapi

import (
	"encoding/json"
	"fmt"
)

type CollectionsToys struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Toys []struct {
		Toy struct {
			Id  float64 `json:"id"`
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
		} `json:"toy"`
	} `json:"toys"`
}

func (req RequestFunc) CharacterCollectionsToys(realm string, name string) (s CollectionsToys, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/collections/toys", realm, name)
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

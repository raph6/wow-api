package wowapi

import (
	"encoding/json"
	"fmt"
)

type Collections struct {
	Heirlooms struct {
		Href string `json:"href"`
	} `json:"heirlooms"`
	Toys struct {
		Href string `json:"href"`
	} `json:"toys"`
	Character struct {
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
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
	} `json:"character"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Pets struct {
		Href string `json:"href"`
	} `json:"pets"`
	Mounts struct {
		Href string `json:"href"`
	} `json:"mounts"`
}

func (req RequestFunc) CharacterCollections(realm string, name string) (s Collections, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/collections", realm, name)
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

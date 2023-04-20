package wowapi

import (
	"encoding/json"
	"fmt"
)

type MythicKeystoneProfile struct {
	CurrentMythicRating struct {
		Color struct {
			R float64 `json:"r"`
			G float64 `json:"g"`
			B float64 `json:"b"`
			A float64 `json:"a"`
		} `json:"color"`
		Rating float64 `json:"rating"`
	} `json:"current_mythic_rating"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	CurrentPeriod struct {
		Period struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Id float64 `json:"id"`
		} `json:"period"`
	} `json:"current_period"`
	Seasons []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Id float64 `json:"id"`
	} `json:"seasons"`
	Character struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name  string  `json:"name"`
		Id    float64 `json:"id"`
		Realm struct {
			Slug string `json:"slug"`
			Key  struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"realm"`
	} `json:"character"`
}

func (req RequestFunc) CharacterMythicKeystoneProfile(realm string, name string) (s MythicKeystoneProfile, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/mythic-keystone-profile", realm, name)
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

package wowapi

import (
	"encoding/json"
	"fmt"
)

type PvpSummary struct {
	HonorableKills float64 `json:"honorable_kills"`
	Character      struct {
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
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Brackets []struct {
		Href string `json:"href"`
	} `json:"brackets"`
	HonorLevel       float64 `json:"honor_level"`
	PvpMapStatistics []struct {
		WorldMap struct {
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"world_map"`
		MatchStatistics struct {
			Played float64 `json:"played"`
			Won    float64 `json:"won"`
			Lost   float64 `json:"lost"`
		} `json:"match_statistics"`
	} `json:"pvp_map_statistics"`
}

func (req RequestFunc) CharacterPvpSummary(realm string, name string) (s PvpSummary, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/pvp-summary", realm, name)
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

package wowapi

import (
	"encoding/json"
	"fmt"
)

type PvpBracket struct {
	SeasonRoundStatistics struct {
		Won    float64 `json:"won"`
		Lost   float64 `json:"lost"`
		Played float64 `json:"played"`
	} `json:"season_round_statistics"`
	WeeklyRoundStatistics struct {
		Lost   float64 `json:"lost"`
		Played float64 `json:"played"`
		Won    float64 `json:"won"`
	} `json:"weekly_round_statistics"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Faction struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"faction"`
	Bracket struct {
		Type string  `json:"type"`
		Id   float64 `json:"id"`
	} `json:"bracket"`
	Season struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Id float64 `json:"id"`
	} `json:"season"`
	Tier struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Id float64 `json:"id"`
	} `json:"tier"`
	WeeklyMatchStatistics struct {
		Won    float64 `json:"won"`
		Lost   float64 `json:"lost"`
		Played float64 `json:"played"`
	} `json:"weekly_match_statistics"`
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
	Rating                float64 `json:"rating"`
	SeasonMatchStatistics struct {
		Played float64 `json:"played"`
		Won    float64 `json:"won"`
		Lost   float64 `json:"lost"`
	} `json:"season_match_statistics"`
	Specialization struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string  `json:"name"`
		Id   float64 `json:"id"`
	} `json:"specialization"`
}

func (req RequestFunc) CharacterPvpBracket(realm string, name string, pvpbracket string) (s PvpBracket, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/pvp-bracket/%s", realm, name, pvpbracket)
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

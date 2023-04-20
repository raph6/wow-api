package wowapi

import (
	"encoding/json"
	"fmt"
)

type EncountersRaids struct {
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
	Expansions []struct {
		Expansion struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"expansion"`
		Instances []struct {
			Instance struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string  `json:"name"`
				Id   float64 `json:"id"`
			} `json:"instance"`
			Modes []struct {
				Difficulty struct {
					Type string `json:"type"`
					Name string `json:"name"`
				} `json:"difficulty"`
				Status struct {
					Type string `json:"type"`
					Name string `json:"name"`
				} `json:"status"`
				Progress struct {
					CompletedCount float64 `json:"completed_count"`
					TotalCount     float64 `json:"total_count"`
					Encounters     []struct {
						Encounter struct {
							Id  float64 `json:"id"`
							Key struct {
								Href string `json:"href"`
							} `json:"key"`
							Name string `json:"name"`
						} `json:"encounter"`
						CompletedCount    float64 `json:"completed_count"`
						LastKillTimestamp float64 `json:"last_kill_timestamp"`
					} `json:"encounters"`
				} `json:"progress"`
			} `json:"modes"`
		} `json:"instances"`
	} `json:"expansions"`
}

func (req RequestFunc) CharacterEncountersRaids(realm string, name string) (s EncountersRaids, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/encounters/raids", realm, name)
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

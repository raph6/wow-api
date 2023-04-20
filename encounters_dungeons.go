package wowapi

import (
	"encoding/json"
	"fmt"
)

type EncountersDungeons struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
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
				Name string  `json:"name"`
				Id   float64 `json:"id"`
				Key  struct {
					Href string `json:"href"`
				} `json:"key"`
			} `json:"instance"`
			Modes []struct {
				Difficulty struct {
					Name string `json:"name"`
					Type string `json:"type"`
				} `json:"difficulty"`
				Status struct {
					Name string `json:"name"`
					Type string `json:"type"`
				} `json:"status"`
				Progress struct {
					Encounters []struct {
						Encounter struct {
							Key struct {
								Href string `json:"href"`
							} `json:"key"`
							Name string  `json:"name"`
							Id   float64 `json:"id"`
						} `json:"encounter"`
						CompletedCount    float64 `json:"completed_count"`
						LastKillTimestamp float64 `json:"last_kill_timestamp"`
					} `json:"encounters"`
					CompletedCount float64 `json:"completed_count"`
					TotalCount     float64 `json:"total_count"`
				} `json:"progress"`
			} `json:"modes"`
		} `json:"instances"`
	} `json:"expansions"`
}

func (req RequestFunc) CharacterEncountersDungeons(realm string, name string) (s EncountersDungeons, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/encounters/dungeons", realm, name)
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

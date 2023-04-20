package wowapi

import (
	"encoding/json"
	"fmt"
)

type AchievementsStatistics struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Character struct {
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
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
	} `json:"character"`
	Categories []struct {
		Id            float64 `json:"id"`
		Name          string  `json:"name"`
		SubCategories []struct {
			Id         float64 `json:"id"`
			Name       string  `json:"name"`
			Statistics []struct {
				Id                   float64 `json:"id"`
				Name                 string  `json:"name"`
				LastUpdatedTimestamp float64 `json:"last_updated_timestamp"`
				Quantity             float64 `json:"quantity"`
			} `json:"statistics"`
		} `json:"sub_categories"`
		Statistics []struct {
			Name                 string  `json:"name"`
			LastUpdatedTimestamp float64 `json:"last_updated_timestamp"`
			Quantity             float64 `json:"quantity"`
			Id                   float64 `json:"id"`
		} `json:"statistics"`
	} `json:"categories"`
}

func (req RequestFunc) CharacterAchievementsStatistics(realm string, name string) (s AchievementsStatistics, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/achievements/statistics", realm, name)
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

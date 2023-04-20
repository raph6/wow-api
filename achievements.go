package wowapi

import (
	"encoding/json"
	"fmt"
)

type Achievements struct {
	RecentEvents []struct {
		Achievement struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"achievement"`
		Timestamp float64 `json:"timestamp"`
	} `json:"recent_events"`
	Character struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name  string  `json:"name"`
		Id    float64 `json:"id"`
		Realm struct {
			Id   float64 `json:"id"`
			Slug string  `json:"slug"`
			Key  struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
		} `json:"realm"`
	} `json:"character"`
	Statistics struct {
		Href string `json:"href"`
	} `json:"statistics"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	TotalQuantity float64 `json:"total_quantity"`
	TotalPoints   float64 `json:"total_points"`
	Achievements  []struct {
		Id          float64 `json:"id"`
		Achievement struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"achievement"`
		Criteria struct {
			Id          float64 `json:"id"`
			IsCompleted bool    `json:"is_completed"`
		} `json:"criteria"`
		CompletedTimestamp float64 `json:"completed_timestamp"`
	} `json:"achievements"`
	CategoryProgress []struct {
		Points   float64 `json:"points"`
		Category struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"category"`
		Quantity float64 `json:"quantity"`
	} `json:"category_progress"`
}

func (req RequestFunc) CharacterAchievements(realm string, name string) (s Achievements, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/achievements", realm, name)
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

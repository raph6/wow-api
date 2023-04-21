package wowapi

import (
	"encoding/json"
	"fmt"
)

type Character struct {
	Realm struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string  `json:"name"`
		Id   float64 `json:"id"`
		Slug string  `json:"slug"`
	} `json:"realm"`
	Specializations struct {
		Href string `json:"href"`
	} `json:"specializations"`
	Statistics struct {
		Href string `json:"href"`
	} `json:"statistics"`
	Gender struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"gender"`
	EquippedItemLevel float64 `json:"equipped_item_level"`
	Experience        float64 `json:"experience"`
	Reputations       struct {
		Href string `json:"href"`
	} `json:"reputations"`
	Professions struct {
		Href string `json:"href"`
	} `json:"professions"`
	AchievementsStatistics struct {
		Href string `json:"href"`
	} `json:"achievements_statistics"`
	Faction struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"faction"`
	CharacterClass struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string  `json:"name"`
		Id   float64 `json:"id"`
	} `json:"character_class"`
	Achievements struct {
		Href string `json:"href"`
	} `json:"achievements"`
	PvpSummary struct {
		Href string `json:"href"`
	} `json:"pvp_summary"`
	Encounters struct {
		Href string `json:"href"`
	} `json:"encounters"`
	Media struct {
		Href string `json:"href"`
	} `json:"media"`
	Equipment struct {
		Href string `json:"href"`
	} `json:"equipment"`
	CovenantProgress struct {
		ChosenCovenant struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"chosen_covenant"`
		RenownLevel float64 `json:"renown_level"`
		Soulbinds   struct {
			Href string `json:"href"`
		} `json:"soulbinds"`
	} `json:"covenant_progress"`
	ActiveSpec struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string  `json:"name"`
		Id   float64 `json:"id"`
	} `json:"active_spec"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Level      float64 `json:"level"`
	Appearance struct {
		Href string `json:"href"`
	} `json:"appearance"`
	Collections struct {
		Href string `json:"href"`
	} `json:"collections"`
	Id   float64 `json:"id"`
	Race struct {
		Id  float64 `json:"id"`
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
	} `json:"race"`
	AchievementPoints float64 `json:"achievement_points"`
	Titles            struct {
		Href string `json:"href"`
	} `json:"titles"`
	LastLoginTimestamp float64 `json:"last_login_timestamp"`
	AverageItemLevel   float64 `json:"average_item_level"`
	Name               string  `json:"name"`
	Guild              struct {
		Faction struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"faction"`
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
	} `json:"guild"`
	MythicKeystoneProfile struct {
		Href string `json:"href"`
	} `json:"mythic_keystone_profile"`
	ActiveTitle struct {
		DisplayString string `json:"display_string"`
		Key           struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string  `json:"name"`
		Id   float64 `json:"id"`
	} `json:"active_title"`
	Quests struct {
		Href string `json:"href"`
	} `json:"quests"`
}

func (req RequestFunc) CharacterInfo(realm string, name string) (s Character, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s", realm, name)
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

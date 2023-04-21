package wowapi

import (
	"encoding/json"
	"fmt"
)

type Professions struct {
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
			Id   float64 `json:"id"`
			Slug string  `json:"slug"`
			Key  struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
		} `json:"realm"`
	} `json:"character"`
	Primaries []struct {
		Profession struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"profession"`
		Tiers []struct {
			KnownRecipes []struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string  `json:"name"`
				Id   float64 `json:"id"`
			} `json:"known_recipes"`
			SkillPoints    float64 `json:"skill_points"`
			MaxSkillPoints float64 `json:"max_skill_points"`
			Tier           struct {
				Id   float64 `json:"id"`
				Name string  `json:"name"`
			} `json:"tier"`
		} `json:"tiers"`
		Specialization struct {
			Name string `json:"name"`
		} `json:"specialization"`
	} `json:"primaries"`
	Secondaries []struct {
		Profession struct {
			Name string  `json:"name"`
			Id   float64 `json:"id"`
			Key  struct {
				Href string `json:"href"`
			} `json:"key"`
		} `json:"profession"`
		SkillPoints    float64 `json:"skill_points"`
		MaxSkillPoints float64 `json:"max_skill_points"`
	} `json:"secondaries"`
}

func (req RequestFunc) CharacterProfessions(realm string, name string) (s Professions, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/professions", realm, name)
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

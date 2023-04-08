package wowapi

import (
	"encoding/json"
	"fmt"
)

type Statistics struct {
	Crit        Rating  `json:"melee_crit"`
	Haste       Rating  `json:"melee_haste"`
	Mastery     Rating  `json:"mastery"`
	Versatility float64 `json:"versatility_damage_done_bonus"`
}

type Rating struct {
	Value float64 `json:"value"`
}

func (req RequestFunc) CharacterStats(realm string, name string) (stats interface{}, err error) {
	url := fmt.Sprintf("https://eu.api.blizzard.com/profile/wow/character/%s/%s/statistics?namespace=profile-eu&locale=fr_FR", realm, name)
	body, err := req(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &stats)
	if err != nil {
		return nil, err
	}
	return stats, nil
}

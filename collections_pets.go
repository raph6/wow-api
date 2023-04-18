package wowapi

type CollectionsPets struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Pets []struct {
		Species struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"species"`
		Level   float64 `json:"level"`
		Quality struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"quality"`
		Stats struct {
			BreedId float64 `json:"breed_id"`
			Health  float64 `json:"health"`
			Power   float64 `json:"power"`
			Speed   float64 `json:"speed"`
		} `json:"stats"`
		CreatureDisplay struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Id float64 `json:"id"`
		} `json:"creature_display"`
		Id float64 `json:"id"`
	} `json:"pets"`
	UnlockedBattlePetSlots float64 `json:"unlocked_battle_pet_slots"`
}

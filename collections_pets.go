package wowapi

type CollectionsPets struct {
	_links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Pets []struct {
		Quality struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"quality"`
		Stats struct {
			Power    float64 `json:"power"`
			Speed    float64 `json:"speed"`
			Breed_id float64 `json:"breed_id"`
			Health   float64 `json:"health"`
		} `json:"stats"`
		Creature_display struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Id float64 `json:"id"`
		} `json:"creature_display"`
		Id      float64 `json:"id"`
		Species struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"species"`
		Level float64 `json:"level"`
	} `json:"pets"`
	Unlocked_battle_pet_slots float64 `json:"unlocked_battle_pet_slots"`
}

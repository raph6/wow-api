package wowapi

type Statistics struct {
	Crit        Rating  `json:"melee_crit"`
	Haste       Rating  `json:"melee_haste"`
	Mastery     Rating  `json:"mastery"`
	Versatility float64 `json:"versatility_damage_done_bonus"`
}

type Rating struct {
	Value float64 `json:"value"`
}

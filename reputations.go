package wowapi

type Reputations struct {
	Reputations []struct {
		Faction struct {
			Name string `json:"name"`
			Id   int    `json:"id"`
		} `json:"faction"`
		Standing struct {
			Name        string `json:"name"`
			RenownLevel int    `json:"renown_level"`
		} `json:"standing"`
	} `json:"reputations"`
}

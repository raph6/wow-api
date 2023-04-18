package wowapi

type PvpSummary struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Brackets []struct {
		Href string `json:"href"`
	} `json:"brackets"`
	HonorLevel       float64 `json:"honor_level"`
	PvpMapStatistics []struct {
		WorldMap struct {
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"world_map"`
		MatchStatistics struct {
			Played float64 `json:"played"`
			Won    float64 `json:"won"`
			Lost   float64 `json:"lost"`
		} `json:"match_statistics"`
	} `json:"pvp_map_statistics"`
	HonorableKills float64 `json:"honorable_kills"`
	Character      struct {
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
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
	} `json:"character"`
}

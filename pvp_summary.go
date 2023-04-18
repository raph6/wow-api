package wowapi

type PvpSummary struct {
	Honor_level        float64 `json:"honor_level"`
	Pvp_map_statistics []struct {
		World_map struct {
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"world_map"`
		Match_statistics struct {
			Played float64 `json:"played"`
			Won    float64 `json:"won"`
			Lost   float64 `json:"lost"`
		} `json:"match_statistics"`
	} `json:"pvp_map_statistics"`
	Honorable_kills float64 `json:"honorable_kills"`
	Character       struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
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
	} `json:"character"`
	_links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Brackets []struct {
		Href string `json:"href"`
	} `json:"brackets"`
}

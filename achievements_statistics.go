package wowapi

type AchievementsStatistics struct {
	_links struct {
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
	Categories []struct {
		Id             float64 `json:"id"`
		Name           string  `json:"name"`
		Sub_categories []struct {
			Statistics []struct {
				Id                     float64 `json:"id"`
				Name                   string  `json:"name"`
				Last_updated_timestamp float64 `json:"last_updated_timestamp"`
				Quantity               float64 `json:"quantity"`
			} `json:"statistics"`
			Id   float64 `json:"id"`
			Name string  `json:"name"`
		} `json:"sub_categories"`
		Statistics []struct {
			Id                     float64 `json:"id"`
			Name                   string  `json:"name"`
			Last_updated_timestamp float64 `json:"last_updated_timestamp"`
			Quantity               float64 `json:"quantity"`
		} `json:"statistics"`
	} `json:"categories"`
}

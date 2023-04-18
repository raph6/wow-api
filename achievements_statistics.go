package wowapi

type AchievementsStatistics struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Character struct {
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
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
	} `json:"character"`
	Categories []struct {
		Statistics []struct {
			Id                   float64 `json:"id"`
			Name                 string  `json:"name"`
			LastUpdatedTimestamp float64 `json:"last_updated_timestamp"`
			Quantity             float64 `json:"quantity"`
		} `json:"statistics"`
		Id            float64 `json:"id"`
		Name          string  `json:"name"`
		SubCategories []struct {
			Id         float64 `json:"id"`
			Name       string  `json:"name"`
			Statistics []struct {
				Id                   float64 `json:"id"`
				Name                 string  `json:"name"`
				LastUpdatedTimestamp float64 `json:"last_updated_timestamp"`
				Quantity             float64 `json:"quantity"`
			} `json:"statistics"`
		} `json:"sub_categories"`
	} `json:"categories"`
}

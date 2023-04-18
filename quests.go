package wowapi

type Quests struct {
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
	InProgress []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string  `json:"name"`
		Id   float64 `json:"id"`
	} `json:"in_progress"`
	Completed struct {
		Href string `json:"href"`
	} `json:"completed"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

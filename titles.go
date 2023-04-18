package wowapi

type Titles struct {
	Titles []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string  `json:"name"`
		Id   float64 `json:"id"`
	} `json:"titles"`
	Links struct {
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
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
			Slug string  `json:"slug"`
		} `json:"realm"`
	} `json:"character"`
	ActiveTitle struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name          string  `json:"name"`
		Id            float64 `json:"id"`
		DisplayString string  `json:"display_string"`
	} `json:"active_title"`
}

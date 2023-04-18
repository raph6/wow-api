package wowapi

type Encounters struct {
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
			Name string  `json:"name"`
			Id   float64 `json:"id"`
			Slug string  `json:"slug"`
			Key  struct {
				Href string `json:"href"`
			} `json:"key"`
		} `json:"realm"`
	} `json:"character"`
	Dungeons struct {
		Href string `json:"href"`
	} `json:"dungeons"`
	Raids struct {
		Href string `json:"href"`
	} `json:"raids"`
}

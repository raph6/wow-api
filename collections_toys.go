package wowapi

type CollectionsToys struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Toys []struct {
		Toy struct {
			Id  float64 `json:"id"`
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
		} `json:"toy"`
	} `json:"toys"`
}

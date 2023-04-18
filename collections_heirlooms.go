package wowapi

type CollectionsHeirlooms struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Heirlooms []struct {
		Heirloom struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"heirloom"`
		Upgrade struct {
			Level float64 `json:"level"`
		} `json:"upgrade"`
	} `json:"heirlooms"`
}

package wowapi

type CollectionsMounts struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Mounts []struct {
		Mount struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"mount"`
		IsUseable bool `json:"is_useable"`
	} `json:"mounts"`
}

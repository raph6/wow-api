package wowapi

type Collections struct {
	Heirlooms struct {
		Href string `json:"href"`
	} `json:"heirlooms"`
	Toys struct {
		Href string `json:"href"`
	} `json:"toys"`
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
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Pets struct {
		Href string `json:"href"`
	} `json:"pets"`
	Mounts struct {
		Href string `json:"href"`
	} `json:"mounts"`
}

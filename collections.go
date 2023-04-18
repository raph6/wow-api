package wowapi

type Collections struct {
	Character struct {
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
		Name string  `json:"name"`
		Id   float64 `json:"id"`
	} `json:"character"`
	_links struct {
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
	Heirlooms struct {
		Href string `json:"href"`
	} `json:"heirlooms"`
	Toys struct {
		Href string `json:"href"`
	} `json:"toys"`
}

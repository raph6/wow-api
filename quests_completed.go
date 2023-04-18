package wowapi

type QuestsCompleted struct {
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
			Name string  `json:"name"`
			Id   float64 `json:"id"`
			Slug string  `json:"slug"`
			Key  struct {
				Href string `json:"href"`
			} `json:"key"`
		} `json:"realm"`
	} `json:"character"`
	Quests []struct {
		Name string  `json:"name"`
		Id   float64 `json:"id"`
		Key  struct {
			Href string `json:"href"`
		} `json:"key"`
	} `json:"quests"`
}

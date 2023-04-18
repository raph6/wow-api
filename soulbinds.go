package wowapi

type Soulbinds struct {
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
	ChosenCovenant struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string  `json:"name"`
		Id   float64 `json:"id"`
	} `json:"chosen_covenant"`
	RenownLevel float64 `json:"renown_level"`
	Soulbinds   []struct {
		Soulbind struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"soulbind"`
		Traits []struct {
			Trait struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string  `json:"name"`
				Id   float64 `json:"id"`
			} `json:"trait"`
			Tier         float64 `json:"tier"`
			DisplayOrder float64 `json:"display_order"`
		} `json:"traits"`
	} `json:"soulbinds"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

package wowapi

type EncountersRaids struct {
	_links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Character struct {
		Id    float64 `json:"id"`
		Realm struct {
			Id   float64 `json:"id"`
			Slug string  `json:"slug"`
			Key  struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
		} `json:"realm"`
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
	} `json:"character"`
	Expansions []struct {
		Expansion struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"expansion"`
		Instances []struct {
			Instance struct {
				Name string  `json:"name"`
				Id   float64 `json:"id"`
				Key  struct {
					Href string `json:"href"`
				} `json:"key"`
			} `json:"instance"`
			Modes []struct {
				Status struct {
					Name string `json:"name"`
					Type string `json:"type"`
				} `json:"status"`
				Progress struct {
					Encounters []struct {
						Last_kill_timestamp float64 `json:"last_kill_timestamp"`
						Encounter           struct {
							Key struct {
								Href string `json:"href"`
							} `json:"key"`
							Name string  `json:"name"`
							Id   float64 `json:"id"`
						} `json:"encounter"`
						Completed_count float64 `json:"completed_count"`
					} `json:"encounters"`
					Completed_count float64 `json:"completed_count"`
					Total_count     float64 `json:"total_count"`
				} `json:"progress"`
				Difficulty struct {
					Type string `json:"type"`
					Name string `json:"name"`
				} `json:"difficulty"`
			} `json:"modes"`
		} `json:"instances"`
	} `json:"expansions"`
}

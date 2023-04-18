package wowapi

type EncountersDungeons struct {
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
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string  `json:"name"`
				Id   float64 `json:"id"`
			} `json:"instance"`
			Modes []struct {
				Difficulty struct {
					Name string `json:"name"`
					Type string `json:"type"`
				} `json:"difficulty"`
				Status struct {
					Type string `json:"type"`
					Name string `json:"name"`
				} `json:"status"`
				Progress struct {
					Encounters []struct {
						Encounter struct {
							Key struct {
								Href string `json:"href"`
							} `json:"key"`
							Name string  `json:"name"`
							Id   float64 `json:"id"`
						} `json:"encounter"`
						Completed_count     float64 `json:"completed_count"`
						Last_kill_timestamp float64 `json:"last_kill_timestamp"`
					} `json:"encounters"`
					Completed_count float64 `json:"completed_count"`
					Total_count     float64 `json:"total_count"`
				} `json:"progress"`
			} `json:"modes"`
		} `json:"instances"`
	} `json:"expansions"`
	_links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

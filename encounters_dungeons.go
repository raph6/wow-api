package wowapi

type EncountersDungeons struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
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
					Type string `json:"type"`
					Name string `json:"name"`
				} `json:"difficulty"`
				Status struct {
					Type string `json:"type"`
					Name string `json:"name"`
				} `json:"status"`
				Progress struct {
					TotalCount float64 `json:"total_count"`
					Encounters []struct {
						Encounter struct {
							Id  float64 `json:"id"`
							Key struct {
								Href string `json:"href"`
							} `json:"key"`
							Name string `json:"name"`
						} `json:"encounter"`
						CompletedCount    float64 `json:"completed_count"`
						LastKillTimestamp float64 `json:"last_kill_timestamp"`
					} `json:"encounters"`
					CompletedCount float64 `json:"completed_count"`
				} `json:"progress"`
			} `json:"modes"`
		} `json:"instances"`
	} `json:"expansions"`
}

package wowapi

type MythicKeystoneProfile struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	CurrentPeriod struct {
		Period struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Id float64 `json:"id"`
		} `json:"period"`
		BestRuns []struct {
			Duration float64 `json:"duration"`
			Members  []struct {
				Character struct {
					Name  string  `json:"name"`
					Id    float64 `json:"id"`
					Realm struct {
						Key struct {
							Href string `json:"href"`
						} `json:"key"`
						Id   float64 `json:"id"`
						Slug string  `json:"slug"`
					} `json:"realm"`
				} `json:"character"`
				Specialization struct {
					Key struct {
						Href string `json:"href"`
					} `json:"key"`
					Name string  `json:"name"`
					Id   float64 `json:"id"`
				} `json:"specialization"`
				Race struct {
					Key struct {
						Href string `json:"href"`
					} `json:"key"`
					Name string  `json:"name"`
					Id   float64 `json:"id"`
				} `json:"race"`
				EquippedItemLevel float64 `json:"equipped_item_level"`
			} `json:"members"`
			IsCompletedWithinTime bool `json:"is_completed_within_time"`
			MythicRating          struct {
				Color struct {
					R float64 `json:"r"`
					G float64 `json:"g"`
					B float64 `json:"b"`
					A float64 `json:"a"`
				} `json:"color"`
				Rating float64 `json:"rating"`
			} `json:"mythic_rating"`
			MapRating struct {
				Color struct {
					G float64 `json:"g"`
					B float64 `json:"b"`
					A float64 `json:"a"`
					R float64 `json:"r"`
				} `json:"color"`
				Rating float64 `json:"rating"`
			} `json:"map_rating"`
			CompletedTimestamp float64 `json:"completed_timestamp"`
			KeystoneLevel      float64 `json:"keystone_level"`
			KeystoneAffixes    []struct {
				Id  float64 `json:"id"`
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
			} `json:"keystone_affixes"`
			Dungeon struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string  `json:"name"`
				Id   float64 `json:"id"`
			} `json:"dungeon"`
		} `json:"best_runs"`
	} `json:"current_period"`
	Seasons []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Id float64 `json:"id"`
	} `json:"seasons"`
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
	CurrentMythicRating struct {
		Color struct {
			R float64 `json:"r"`
			G float64 `json:"g"`
			B float64 `json:"b"`
			A float64 `json:"a"`
		} `json:"color"`
		Rating float64 `json:"rating"`
	} `json:"current_mythic_rating"`
}

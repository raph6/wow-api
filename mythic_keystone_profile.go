package wowapi

type MythicKeystoneProfile struct {
	Current_mythic_rating struct {
		Color struct {
			R float64 `json:"r"`
			G float64 `json:"g"`
			B float64 `json:"b"`
			A float64 `json:"a"`
		} `json:"color"`
		Rating float64 `json:"rating"`
	} `json:"current_mythic_rating"`
	_links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Current_period struct {
		Period struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Id float64 `json:"id"`
		} `json:"period"`
		Best_runs []struct {
			Completed_timestamp float64 `json:"completed_timestamp"`
			Keystone_level      float64 `json:"keystone_level"`
			Keystone_affixes    []struct {
				Name string  `json:"name"`
				Id   float64 `json:"id"`
				Key  struct {
					Href string `json:"href"`
				} `json:"key"`
			} `json:"keystone_affixes"`
			Dungeon struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string  `json:"name"`
				Id   float64 `json:"id"`
			} `json:"dungeon"`
			Mythic_rating struct {
				Color struct {
					R float64 `json:"r"`
					G float64 `json:"g"`
					B float64 `json:"b"`
					A float64 `json:"a"`
				} `json:"color"`
				Rating float64 `json:"rating"`
			} `json:"mythic_rating"`
			Duration float64 `json:"duration"`
			Members  []struct {
				Equipped_item_level float64 `json:"equipped_item_level"`
				Character           struct {
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
					Name string  `json:"name"`
					Id   float64 `json:"id"`
					Key  struct {
						Href string `json:"href"`
					} `json:"key"`
				} `json:"race"`
			} `json:"members"`
			Is_completed_within_time bool `json:"is_completed_within_time"`
			Map_rating               struct {
				Color struct {
					B float64 `json:"b"`
					A float64 `json:"a"`
					R float64 `json:"r"`
					G float64 `json:"g"`
				} `json:"color"`
				Rating float64 `json:"rating"`
			} `json:"map_rating"`
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
			Slug string `json:"slug"`
			Key  struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"realm"`
	} `json:"character"`
}

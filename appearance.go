package wowapi

type Appearance struct {
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
	Playable_class struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string  `json:"name"`
		Id   float64 `json:"id"`
	} `json:"playable_class"`
	Active_spec struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string  `json:"name"`
		Id   float64 `json:"id"`
	} `json:"active_spec"`
	Gender struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"gender"`
	Customizations []struct {
		Option struct {
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"option"`
		Choice struct {
			Id            float64 `json:"id"`
			Display_order float64 `json:"display_order"`
		} `json:"choice"`
	} `json:"customizations"`
	_links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Playable_race struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string  `json:"name"`
		Id   float64 `json:"id"`
	} `json:"playable_race"`
	Faction struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"faction"`
	Guild_crest struct {
		Emblem struct {
			Id    float64 `json:"id"`
			Media struct {
				Id  float64 `json:"id"`
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
			} `json:"media"`
			Color struct {
				Id   float64 `json:"id"`
				Rgba struct {
					A float64 `json:"a"`
					R float64 `json:"r"`
					G float64 `json:"g"`
					B float64 `json:"b"`
				} `json:"rgba"`
			} `json:"color"`
		} `json:"emblem"`
		Border struct {
			Id    float64 `json:"id"`
			Media struct {
				Id  float64 `json:"id"`
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
			} `json:"media"`
			Color struct {
				Rgba struct {
					B float64 `json:"b"`
					A float64 `json:"a"`
					R float64 `json:"r"`
					G float64 `json:"g"`
				} `json:"rgba"`
				Id float64 `json:"id"`
			} `json:"color"`
		} `json:"border"`
		Background struct {
			Color struct {
				Id   float64 `json:"id"`
				Rgba struct {
					R float64 `json:"r"`
					G float64 `json:"g"`
					B float64 `json:"b"`
					A float64 `json:"a"`
				} `json:"rgba"`
			} `json:"color"`
		} `json:"background"`
	} `json:"guild_crest"`
	Items []struct {
		Id   float64 `json:"id"`
		Slot struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"slot"`
		Enchant                     float64 `json:"enchant"`
		Item_appearance_modifier_id float64 `json:"item_appearance_modifier_id"`
		Internal_slot_id            float64 `json:"internal_slot_id"`
		Subclass                    float64 `json:"subclass"`
	} `json:"items"`
}

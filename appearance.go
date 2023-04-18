package wowapi

type Appearance struct {
	ActiveSpec struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string  `json:"name"`
		Id   float64 `json:"id"`
	} `json:"active_spec"`
	Faction struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"faction"`
	Items []struct {
		Slot struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"slot"`
		Enchant                  float64 `json:"enchant"`
		ItemAppearanceModifierId float64 `json:"item_appearance_modifier_id"`
		InternalSlotId           float64 `json:"internal_slot_id"`
		Subclass                 float64 `json:"subclass"`
		Id                       float64 `json:"id"`
	} `json:"items"`
	Links struct {
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
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
			Slug string  `json:"slug"`
		} `json:"realm"`
	} `json:"character"`
	PlayableRace struct {
		Name string  `json:"name"`
		Id   float64 `json:"id"`
		Key  struct {
			Href string `json:"href"`
		} `json:"key"`
	} `json:"playable_race"`
	Customizations []struct {
		Option struct {
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"option"`
		Choice struct {
			Id           float64 `json:"id"`
			DisplayOrder float64 `json:"display_order"`
		} `json:"choice"`
	} `json:"customizations"`
	PlayableClass struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string  `json:"name"`
		Id   float64 `json:"id"`
	} `json:"playable_class"`
	Gender struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"gender"`
	GuildCrest struct {
		Emblem struct {
			Id    float64 `json:"id"`
			Media struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Id float64 `json:"id"`
			} `json:"media"`
			Color struct {
				Id   float64 `json:"id"`
				Rgba struct {
					R float64 `json:"r"`
					G float64 `json:"g"`
					B float64 `json:"b"`
					A float64 `json:"a"`
				} `json:"rgba"`
			} `json:"color"`
		} `json:"emblem"`
		Border struct {
			Id    float64 `json:"id"`
			Media struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Id float64 `json:"id"`
			} `json:"media"`
			Color struct {
				Id   float64 `json:"id"`
				Rgba struct {
					R float64 `json:"r"`
					G float64 `json:"g"`
					B float64 `json:"b"`
					A float64 `json:"a"`
				} `json:"rgba"`
			} `json:"color"`
		} `json:"border"`
		Background struct {
			Color struct {
				Id   float64 `json:"id"`
				Rgba struct {
					A float64 `json:"a"`
					R float64 `json:"r"`
					G float64 `json:"g"`
					B float64 `json:"b"`
				} `json:"rgba"`
			} `json:"color"`
		} `json:"background"`
	} `json:"guild_crest"`
}

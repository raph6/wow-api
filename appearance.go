package wowapi

import (
	"encoding/json"
	"fmt"
)

type Appearance struct {
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
	PlayableRace struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string  `json:"name"`
		Id   float64 `json:"id"`
	} `json:"playable_race"`
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
					B float64 `json:"b"`
					A float64 `json:"a"`
					R float64 `json:"r"`
					G float64 `json:"g"`
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
		Subclass float64 `json:"subclass"`
		Id       float64 `json:"id"`
		Slot     struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"slot"`
		Enchant                  float64 `json:"enchant"`
		ItemAppearanceModifierId float64 `json:"item_appearance_modifier_id"`
		InternalSlotId           float64 `json:"internal_slot_id"`
	} `json:"items"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
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
}

func (req RequestFunc) CharacterAppearance(realm string, name string) (s Appearance, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/appearance", realm, name)
	body, err := req(url)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &s)
	if err != nil {
		return
	}
	return
}

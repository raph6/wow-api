package wowapi

import (
	"encoding/json"
	"fmt"
)

type Specializations struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Specializations []struct {
		Specialization struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"specialization"`
		Glyphs []struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"glyphs"`
		PvpTalentSlots []struct {
			SlotNumber float64 `json:"slot_number"`
			Selected   struct {
				Talent struct {
					Name string  `json:"name"`
					Id   float64 `json:"id"`
					Key  struct {
						Href string `json:"href"`
					} `json:"key"`
				} `json:"talent"`
				SpellTooltip struct {
					Spell struct {
						Key struct {
							Href string `json:"href"`
						} `json:"key"`
						Name string  `json:"name"`
						Id   float64 `json:"id"`
					} `json:"spell"`
					Description string `json:"description"`
					CastTime    string `json:"cast_time"`
					PowerCost   string `json:"power_cost"`
					Range       string `json:"range"`
					Cooldown    string `json:"cooldown"`
				} `json:"spell_tooltip"`
			} `json:"selected"`
		} `json:"pvp_talent_slots"`
		Loadouts []struct {
			IsActive             bool   `json:"is_active"`
			TalentLoadoutCode    string `json:"talent_loadout_code"`
			SelectedClassTalents []struct {
				Id      float64 `json:"id"`
				Rank    float64 `json:"rank"`
				Tooltip struct {
					Talent struct {
						Key struct {
							Href string `json:"href"`
						} `json:"key"`
						Name string  `json:"name"`
						Id   float64 `json:"id"`
					} `json:"talent"`
					SpellTooltip struct {
						PowerCost string `json:"power_cost"`
						Cooldown  string `json:"cooldown"`
						Spell     struct {
							Key struct {
								Href string `json:"href"`
							} `json:"key"`
							Name string  `json:"name"`
							Id   float64 `json:"id"`
						} `json:"spell"`
						Description string `json:"description"`
						CastTime    string `json:"cast_time"`
					} `json:"spell_tooltip"`
				} `json:"tooltip"`
				DefaultPoints float64 `json:"default_points"`
			} `json:"selected_class_talents"`
			SelectedSpecTalents []struct {
				Tooltip struct {
					Talent struct {
						Key struct {
							Href string `json:"href"`
						} `json:"key"`
						Name string  `json:"name"`
						Id   float64 `json:"id"`
					} `json:"talent"`
					SpellTooltip struct {
						Description string `json:"description"`
						CastTime    string `json:"cast_time"`
						PowerCost   string `json:"power_cost"`
						Range       string `json:"range"`
						Spell       struct {
							Key struct {
								Href string `json:"href"`
							} `json:"key"`
							Name string  `json:"name"`
							Id   float64 `json:"id"`
						} `json:"spell"`
					} `json:"spell_tooltip"`
				} `json:"tooltip"`
				Id   float64 `json:"id"`
				Rank float64 `json:"rank"`
			} `json:"selected_spec_talents"`
		} `json:"loadouts"`
	} `json:"specializations"`
	ActiveSpecialization struct {
		Name string  `json:"name"`
		Id   float64 `json:"id"`
		Key  struct {
			Href string `json:"href"`
		} `json:"key"`
	} `json:"active_specialization"`
	Character struct {
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
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
	} `json:"character"`
}

func (req RequestFunc) CharacterSpecializations(realm string, name string) (s Specializations, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/specializations", realm, name)
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

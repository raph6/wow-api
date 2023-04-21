package wowapi

import (
	"encoding/json"
	"fmt"
)

type Statistics struct {
	SpellCrit struct {
		Rating      float64 `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
		Value       float64 `json:"value"`
	} `json:"spell_crit"`
	ManaRegenCombat            float64 `json:"mana_regen_combat"`
	Power                      float64 `json:"power"`
	VersatilityDamageDoneBonus float64 `json:"versatility_damage_done_bonus"`
	MainHandDps                float64 `json:"main_hand_dps"`
	OffHandDps                 float64 `json:"off_hand_dps"`
	Block                      struct {
		Rating      float64 `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
		Value       float64 `json:"value"`
	} `json:"block"`
	RangedHaste struct {
		Rating      float64 `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
		Value       float64 `json:"value"`
	} `json:"ranged_haste"`
	PowerType struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string  `json:"name"`
		Id   float64 `json:"id"`
	} `json:"power_type"`
	Lifesteal struct {
		Rating      float64 `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
		Value       float64 `json:"value"`
	} `json:"lifesteal"`
	VersatilityHealingDoneBonus float64 `json:"versatility_healing_done_bonus"`
	MainHandDamageMax           float64 `json:"main_hand_damage_max"`
	Avoidance                   struct {
		Rating      float64 `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
	} `json:"avoidance"`
	ManaRegen float64 `json:"mana_regen"`
	Health    float64 `json:"health"`
	Strength  struct {
		Base      float64 `json:"base"`
		Effective float64 `json:"effective"`
	} `json:"strength"`
	Intellect struct {
		Base      float64 `json:"base"`
		Effective float64 `json:"effective"`
	} `json:"intellect"`
	VersatilityDamageTakenBonus float64 `json:"versatility_damage_taken_bonus"`
	Dodge                       struct {
		Rating      float64 `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
		Value       float64 `json:"value"`
	} `json:"dodge"`
	Speed struct {
		RatingBonus float64 `json:"rating_bonus"`
		Rating      float64 `json:"rating"`
	} `json:"speed"`
	MeleeCrit struct {
		Rating      float64 `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
		Value       float64 `json:"value"`
	} `json:"melee_crit"`
	MainHandSpeed    float64 `json:"main_hand_speed"`
	OffHandDamageMin float64 `json:"off_hand_damage_min"`
	OffHandSpeed     float64 `json:"off_hand_speed"`
	SpellPower       float64 `json:"spell_power"`
	RangedCrit       struct {
		Rating      float64 `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
		Value       float64 `json:"value"`
	} `json:"ranged_crit"`
	MainHandDamageMin float64 `json:"main_hand_damage_min"`
	OffHandDamageMax  float64 `json:"off_hand_damage_max"`
	SpellPenetration  float64 `json:"spell_penetration"`
	Parry             struct {
		Rating      float64 `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
		Value       float64 `json:"value"`
	} `json:"parry"`
	MeleeHaste struct {
		Rating      float64 `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
		Value       float64 `json:"value"`
	} `json:"melee_haste"`
	Mastery struct {
		Rating      float64 `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
		Value       float64 `json:"value"`
	} `json:"mastery"`
	BonusArmor  float64 `json:"bonus_armor"`
	AttackPower float64 `json:"attack_power"`
	Character   struct {
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
	Agility struct {
		Base      float64 `json:"base"`
		Effective float64 `json:"effective"`
	} `json:"agility"`
	SpellHaste struct {
		Value       float64 `json:"value"`
		Rating      float64 `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
	} `json:"spell_haste"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Stamina struct {
		Base      float64 `json:"base"`
		Effective float64 `json:"effective"`
	} `json:"stamina"`
	Versatility float64 `json:"versatility"`
	Armor       struct {
		Base      float64 `json:"base"`
		Effective float64 `json:"effective"`
	} `json:"armor"`
}

func (req RequestFunc) CharacterStatistics(realm string, name string) (s Statistics, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/statistics", realm, name)
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

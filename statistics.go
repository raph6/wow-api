package wowapi

import (
	"encoding/json"
	"fmt"
)

type Statistics struct {
	Health    int64 `json:"health"`
	Power     int64 `json:"power"`
	PowerType struct {
		Name string `json:"name"`
		Id   int64  `json:"id"`
	} `json:"power_type"`
	Speed                       float64     `json:"speed"`
	Strength                    PrimaryStat `json:"strength"`
	Agility                     PrimaryStat `json:"agility"`
	Intellect                   PrimaryStat `json:"intellect"`
	Stamina                     PrimaryStat `json:"stamina"`
	MeleeCrit                   Rating      `json:"melee_crit"`
	MeleeHaste                  Rating      `json:"melee_haste"`
	Mastery                     Rating      `json:"mastery"`
	BonusArmor                  float64     `json:"bonus_armor"`
	Lifesteal                   Rating      `json:"lifesteal"`
	Versatility                 int64       `json:"versatility"`
	VersatilityDamageDoneBonus  float64     `json:"versatility_damage_done_bonus"`
	VersatilityHealingDoneBonus float64     `json:"versatility_healing_done_bonus"`
	VersatilityDamageTakenBonus float64     `json:"versatility_damage_taken_bonus"`
	Avoidance                   Rating      `json:"avoidance"`
	AttackPower                 int64       `json:"attack_power"`
	MainHandDamageMin           float64     `json:"main_hand_damage_min"`
	MainHandDamageMax           float64     `json:"main_hand_damage_max"`
	MainHandSpeed               float64     `json:"main_hand_speed"`
	MainHandDps                 float64     `json:"main_hand_dps"`
	OffHandDamageMin            float64     `json:"off_hand_damage_min"`
	OffHandDamageMax            float64     `json:"off_hand_damage_max"`
	OffHandSpeed                float64     `json:"off_hand_speed"`
	OffHandDps                  float64     `json:"off_hand_dps"`
	SpellPower                  int64       `json:"spell_power"`
	SpellPenetration            int64       `json:"spell_penetration"`
	SpellCrit                   Rating      `json:"spell_crit"`
	ManaRegen                   float64     `json:"mana_regen"`
	ManaRegenCombat             float64     `json:"mana_regen_combat"`
	Armor                       PrimaryStat `json:"armor"`
	Dodge                       Rating      `json:"dodge"`
	Parry                       Rating      `json:"parry"`
	Block                       Rating      `json:"block"`
	RangedCrit                  Rating      `json:"ranged_crit"`
	RangedHaste                 Rating      `json:"ranged_haste"`
	SpellHaste                  Rating      `json:"spell_haste"`
}

type PrimaryStat struct {
	Base      int64 `json:"base"`
	Effective int64 `json:"effective"`
}

type Rating struct {
	Rating      int     `json:"rating"`
	RatingBonus float64 `json:"rating_bonus"`
	Value       float64 `json:"value"`
}

func (req RequestFunc) CharacterStats(realm string, name string) (stats interface{}, err error) {
	url := fmt.Sprintf("https://eu.api.blizzard.com/profile/wow/character/%s/%s/statistics?namespace=profile-eu&locale=fr_FR", realm, name)
	body, err := req(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &stats)
	if err != nil {
		return nil, err
	}
	return stats, nil
}

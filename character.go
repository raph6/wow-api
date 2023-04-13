package wowapi

type Character struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Gender struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"gender"`
	Faction struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"faction"`
	Race struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"race"`
	CharacterClass struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"character_class"`
	ActiveSpec struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"active_spec"`
	Realm struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
		Slug string `json:"slug"`
	} `json:"realm"`
	Guild struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name  string `json:"name"`
		ID    int    `json:"id"`
		Realm struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
			Slug string `json:"slug"`
		} `json:"realm"`
		Faction struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"faction"`
	} `json:"guild"`
	Level                  int   `json:"level"`
	Experience             int   `json:"experience"`
	AchievementPoints      int   `json:"achievement_points"`
	Achievements           Href  `json:"achievements"`
	Titles                 Href  `json:"titles"`
	PvpSummary             Href  `json:"pvp_summary"`
	Encounters             Href  `json:"encounters"`
	Media                  Href  `json:"media"`
	LastLoginTimestamp     int64 `json:"last_login_timestamp"`
	AverageItemLevel       int   `json:"average_item_level"`
	EquippedItemLevel      int   `json:"equipped_item_level"`
	Specializations        Href  `json:"specializations"`
	Statistics             Href  `json:"statistics"`
	MythicKeystoneProfile  Href  `json:"mythic_keystone_profile"`
	Equipment              Href  `json:"equipment"`
	Appearance             Href  `json:"appearance"`
	Collections            Href  `json:"collections"`
	Reputations            Href  `json:"reputations"`
	Quests                 Href  `json:"quests"`
	AchievementsStatistics Href  `json:"achievements_statistics"`
	Professions            Href  `json:"professions"`
	CovenantProgress       struct {
		ChosenCovenant struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"chosen_covenant"`
		RenownLevel int `json:"renown_level"`
		Soulbinds   struct {
			Href string `json:"href"`
		} `json:"soulbinds"`
	} `json:"covenant_progress"`
}

type Href struct {
	Href string `json:"href"`
}

type Key struct {
	Key struct {
		Href string `json:"href"`
	} `json:"key"`
	Name string `json:"name"`
	ID   int    `json:"id"`
}

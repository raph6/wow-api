package wowapi

type Professions struct {
	Primaries   []Profession `json:"primaries"`
	Secondaries []Profession `json:"secondaries"`
}

type Profession struct {
	Profession struct {
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"profession"`
	Tiers []struct {
		SkillPoints    int `json:"skill_points"`
		MaxSkillPoints int `json:"max_skill_points"`
		Tier           struct {
			Name string `json:"name"`
			Id   int    `json:"id"`
		}
		KnownRecipes []struct {
			Name string `json:"name"`
			Id   int    `json:"id"`
		} `json:"known_recipes"`
	} `json:"tiers"`
}

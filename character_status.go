package wowapi

type CharacterStatus struct {
	IsValid bool `json:"is_valid"`
	Links   struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Id float64 `json:"id"`
}

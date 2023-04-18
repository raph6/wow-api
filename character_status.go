package wowapi

type CharacterStatus struct {
	_links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Id       float64 `json:"id"`
	Is_valid bool    `json:"is_valid"`
}

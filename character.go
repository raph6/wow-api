package wowapi

type Character struct {
	Level              int   `json:"level"`
	ItemLevel          int   `json:"equipped_item_level"`
	BagLevel           int   `json:"average_item_level"`
	LastLoginTimestamp int64 `json:"last_login_timestamp"`
	Guild              struct {
		Name string `json:"name"`
	} `json:"guild"`
}

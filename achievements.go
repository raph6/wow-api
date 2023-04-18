package wowapi

type Achievements struct {
	TotalQuantity float64 `json:"total_quantity"`
	TotalPoints   float64 `json:"total_points"`
	Achievements  []struct {
		Id          float64 `json:"id"`
		Achievement struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"achievement"`
		Criteria struct {
			Id          float64 `json:"id"`
			IsCompleted bool    `json:"is_completed"`
		} `json:"criteria"`
		CompletedTimestamp float64 `json:"completed_timestamp"`
	} `json:"achievements"`
	CategoryProgress []struct {
		Category struct {
			Id  float64 `json:"id"`
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
		} `json:"category"`
		Quantity float64 `json:"quantity"`
		Points   float64 `json:"points"`
	} `json:"category_progress"`
	RecentEvents []struct {
		Timestamp   float64 `json:"timestamp"`
		Achievement struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"achievement"`
	} `json:"recent_events"`
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
	Statistics struct {
		Href string `json:"href"`
	} `json:"statistics"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

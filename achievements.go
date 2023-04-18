package wowapi

type Achievements struct {
	Total_quantity float64 `json:"total_quantity"`
	Total_points   float64 `json:"total_points"`
	Achievements   []struct {
		Id          float64 `json:"id"`
		Achievement struct {
			Name string  `json:"name"`
			Id   float64 `json:"id"`
			Key  struct {
				Href string `json:"href"`
			} `json:"key"`
		} `json:"achievement"`
		Criteria struct {
			Id           float64 `json:"id"`
			Is_completed bool    `json:"is_completed"`
		} `json:"criteria"`
		Completed_timestamp float64 `json:"completed_timestamp"`
	} `json:"achievements"`
	Category_progress []struct {
		Points   float64 `json:"points"`
		Category struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"category"`
		Quantity float64 `json:"quantity"`
	} `json:"category_progress"`
	Recent_events []struct {
		Achievement struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"achievement"`
		Timestamp float64 `json:"timestamp"`
	} `json:"recent_events"`
	Character struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name  string  `json:"name"`
		Id    float64 `json:"id"`
		Realm struct {
			Name string  `json:"name"`
			Id   float64 `json:"id"`
			Slug string  `json:"slug"`
			Key  struct {
				Href string `json:"href"`
			} `json:"key"`
		} `json:"realm"`
	} `json:"character"`
	Statistics struct {
		Href string `json:"href"`
	} `json:"statistics"`
	_links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

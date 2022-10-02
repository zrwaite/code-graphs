package models

type TimeData struct {
	Name         string  `json:"name"`
	Percent      float64 `json:"percent"`
	TotalSeconds float64 `json:"total_seconds"`
	Text         string  `json:"text"`
}

type WakatimeData struct {
	Data struct {
		BestDay struct {
			Date         string `json:"date"`
			TotalSeconds int    `json:"total_seconds"`
			Text         string `json:"text"`
		}
		Categories []struct {
			DailySums map[string]float64 `json:"daily_sums"`
			Name      string             `json:"name"` // "Coding"
		}
		Range            string     `json:"human_readable_range"`
		Status           string     `json:"status"` // "pending_update" or "ok"
		Languages        []TimeData `json:"languages"`
		Editors          []TimeData `json:"editors"`
		OperatingSystems []TimeData `json:"operating_systems"`
		Projects         []TimeData `json:"projects"`
		Username         string     `json:"username"`
		TotalSeconds     float64    `json:"total_seconds"`
	} `json:"data"`
	Expires      string `json:"expires"`
	LastModified string `json:"last_modified"`
}

type WakaTimeRefreshResponse struct {
	AccessToken string `json:"access_token"`
}

package models

import (
	"errors"
	"strings"
)

type TimeData struct {
	Name         string  `json:"name"`
	Percent      float64 `json:"percent"`
	TotalSeconds float64 `json:"total_seconds"`
	Text         string  `json:"text"`
}

type WakatimeStatsData struct {
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
}

type WakatimeData struct {
	Data         WakatimeStatsData `json:"data"`
	Expires      string            `json:"expires"`
	LastModified string            `json:"last_modified"`
	Verified     bool              `json:"verified"`
}

type WakaTimeTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// function on WakaTimeTokenResponse called parseFromString
func (w *WakaTimeTokenResponse) ParseFromString(data string) (err error) {
	// Split the input string by "&"
	pairs := strings.Split(data, "&")

	// Loop through the pairs
	for _, pair := range pairs {
		// Split each pair by "="
		fields := strings.Split(pair, "=")

		// Check the first field and set the appropriate value in the Tokens struct
		switch fields[0] {
		case "access_token":
			w.AccessToken = fields[1]
		case "refresh_token":
			w.RefreshToken = fields[1]
		}
	}

	if w.AccessToken == "" || w.RefreshToken == "" {
		// mail.ErrorMessage("Failed to refresh wakatime token for " + name + " - ")
		return errors.New("failed to find access_token and refresh_token in response")
	}
	return nil
}

type WakaTimeUserResponse struct {
	Data struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"data"`
}

type User struct {
	Username     string `bson:"username"`
	AccessToken  string `bson:"access_token"`
	RefreshToken string `bson:"refresh_token"`
	Verified     bool   `bson:"verified"`
	Public       bool   `bson:"public"`
}

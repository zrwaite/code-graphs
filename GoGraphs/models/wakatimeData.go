package models

type WakatimeData struct {
	Data []struct {
		Languages []struct {
			Name         string  `json:"name"`
			TotalSeconds float64 `json:"total_seconds"`
		} `json:"languages"`
	} `json:"data"`
}

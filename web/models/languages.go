package models

type Language struct {
	Name         string  `json:"name"`
	Colour       string  `json:"colour"`
	TotalSeconds float64 `json:"total_seconds"`
	Percent      float64 `json:"percent"`
}

type Languages struct {
	TotalSeconds float64    `json:"total_seconds"`
	Languages    []Language `json:"languages"`
	Other        struct {
		Languages    []Language `json:"other_languages"`
		TotalSeconds float64    `json:"total_seconds"`
		Percent      float64    `json:"percent"`
	}
}

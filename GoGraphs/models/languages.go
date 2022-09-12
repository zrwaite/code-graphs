package models

type Language struct {
	Name    string  `json:"name"`
	Colour  string  `json:"colour"`
	Time    int     `json:"time"`
	Percent float64 `json:"percent"`
}

type Languages struct {
	TotalTime int        `json:"time"`
	Languages []Language `json:"languages"`
}

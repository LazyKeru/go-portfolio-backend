package models

type Date struct {
	Day   float64 `json:"day"`
	Month float64 `json:"month"`
	Year  float64 `json:"year"`
}

type Tag struct {
	Icon string `json:"icon"`
	Text string `json:"text"`
}

type Experience struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Start       Date   `json:"start"`
	Description string `json:"description"`
	Tags        []Tag  `json:"tags"`
}

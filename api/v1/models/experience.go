package models

type Experience struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Start       Date   `json:"start"`
	Description string `json:"description"`
	Tags        []Tag  `json:"tags"`
}

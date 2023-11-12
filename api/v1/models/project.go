package models

type Project struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:""`
	Tags        []Tag  `json:"tags"`
}

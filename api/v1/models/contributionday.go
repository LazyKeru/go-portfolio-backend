package models

type ContributionDay struct {
	Color             string `json:"color"`
	ContributionCount int    `json:"contributionCount"`
	Date              string `json:"date"`
	Weekday           int    `json:"weekday"`
}

package models

type ContributionCalendar struct {
	Colors             []string           `json:"colors"`
	TotalContributions int                `json:"totalContributions"`
	Weeks              []ContributionDays `json:"weeks"`
}

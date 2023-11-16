package models

type Viewer struct {
	Name                    string                  `json:"name"`
	ContributionsCollection ContributionsCollection `json:"contributionsCollection"`
}

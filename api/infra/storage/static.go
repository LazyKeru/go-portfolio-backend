package storage

import (
	"golang-rest-api-portfolio/api/v1/models"
)

var Static_experiences = []models.Experience{
	{ID: "1", Title: "DevOps - Claranet", Start: models.Date{Day: 10, Month: 12, Year: 2022}, Description: "Stage Claranet", Tags: []models.Tag{{Icon: "pi pi-cloud", Text: "Azure"}, {Icon: "pi pi-box", Text: "k8s"}}},
	{ID: "2", Title: "DevSecOps - Colas", Start: models.Date{Day: 10, Month: 12, Year: 2022}, Description: "Stage Colas", Tags: []models.Tag{{Icon: "pi pi-bolt", Text: "DevOps"}, {Icon: "pi pi-box", Text: "k8s"}}},
}

func GetExperienceByID(id string) *models.Experience {
	var result *models.Experience = nil

	for _, a := range Static_experiences {
		if a.ID == id {
			result = &a
		}
	}

	return result
}

func GetExperiences() []models.Experience {
	return Static_experiences
}

func AddExperience(experience models.Experience) {
	Static_experiences = append(Static_experiences, experience)
}

func RemoveExperience() {

}

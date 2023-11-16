package handlers

import (
	"encoding/json"
	"golang-rest-api-portfolio/api/infra/service/github"
	"golang-rest-api-portfolio/api/v1/models"
	"log"
	"net/http"
)

func GetContribution(w http.ResponseWriter, r *http.Request) {
	res, err := github.GetContribution()
	if err != nil {
		log.Fatal(err)
	}
	contribution := models.ContributionResponse{}
	if err := json.Unmarshal(res, &contribution); err != nil {
		log.Fatal(err)
	}

	response, err := json.Marshal(contribution)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))

}

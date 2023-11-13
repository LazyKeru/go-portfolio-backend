package handlers

import (
	"encoding/json"
	"golang-rest-api-portfolio/api/infra/storage/experience"
	"golang-rest-api-portfolio/api/v1/models"
	"net/http"
)

func GetExperiences(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(experience.GetExperiences())

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func PostExperience(w http.ResponseWriter, r *http.Request) {
	var newExperience models.Experience

	if err := json.NewDecoder(r.Body).Decode(&newExperience); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := json.Marshal(newExperience)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	experience.AddExperience(newExperience)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(res)) // Experience
}

func GetExperienceByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path

	if experience := experience.GetExperienceByID(id); experience == nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		res, err := json.Marshal(experience)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.Write([]byte(res))
	}
}

func DeleteExperience(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path

	if deleted_id := experience.RemoveExperience(id); deleted_id == "" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusNoContent)
		res, err := json.Marshal(deleted_id)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.Write([]byte(res))
	}
}

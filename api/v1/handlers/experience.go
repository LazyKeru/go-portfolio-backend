package handlers

import (
	"encoding/json"
	"golang-rest-api-portfolio/api/infra/storage"
	"golang-rest-api-portfolio/api/v1/models"
	"net/http"
)

func GetExperiences(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(storage.GetExperiences())

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
		return
	}

	storage.AddExperience(newExperience)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Post JSON")) // Experience
}

func GetExperienceByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path

	if experience := storage.GetExperienceByID(id); experience == nil {
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

	if deleted_id := storage.RemoveExperience(id); deleted_id == "" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		res, err := json.Marshal(deleted_id)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.Write([]byte(res))
	}
}

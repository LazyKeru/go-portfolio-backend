package handlers

import (
	"encoding/json"
	"golang-rest-api-portfolio/api/infra/storage/project"
	"golang-rest-api-portfolio/api/v1/models"
	"net/http"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(project.GetProjects())

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func PostProject(w http.ResponseWriter, r *http.Request) {
	var newProject models.Project

	if err := json.NewDecoder(r.Body).Decode(&newProject); err != nil {
		return
	}

	project.AddProject(newProject)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Post JSON")) // Project
}

func GetProjectByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path

	if project := project.GetProjectByID(id); project == nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		res, err := json.Marshal(project)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.Write([]byte(res))
	}
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path

	if deleted_id := project.RemoveProject(id); deleted_id == "" {
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

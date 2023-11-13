package routers

import (
	"golang-rest-api-portfolio/api/v1/handlers"
	"net/http"
)

func SetupProject(prefix string) {
	http.HandleFunc(prefix+"/projects",
		func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "GET":
				handlers.GetProjects(w, r)
			case "POST":
				handlers.PostProject(w, r)
			}

		})
	http.Handle(prefix+"/projects/",
		http.StripPrefix(prefix+"/projects/", http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				// TO DO: Check if the URL.Path is valid
				switch r.Method {
				case "GET":
					handlers.GetProjectByID(w, r)
				case "DELETE":
					handlers.DeleteProject(w, r)
				}
			})))
}

package routers

import (
	"net/http"

	"golang-rest-api-portfolio/api/v1/handlers"
	"golang-rest-api-portfolio/api/v1/middleware"
)

func SetupExperience(prefix string) {
	http.HandleFunc(prefix+"/experiences",
		middleware.JsonHeaderMiddleware(
			func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case "GET":
					handlers.GetExperiences(w, r)
				case "POST":
					handlers.PostExperience(w, r)
				}

			}))
	http.Handle(prefix+"/experiences/",
		http.StripPrefix(prefix+"/experiences/", http.HandlerFunc(
			middleware.JsonHeaderMiddleware(
				func(w http.ResponseWriter, r *http.Request) {
					// TO DO: Check if the URL.Path is valid
					switch r.Method {
					case "GET":
						handlers.GetExperienceByID(w, r)
					case "DELETE":
						handlers.DeleteExperience(w, r)
					}
				}))))
}

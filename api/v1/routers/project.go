package routers

import (
	"golang-rest-api-portfolio/api/v1/handlers"
	"golang-rest-api-portfolio/api/v1/middleware"
	"net/http"
)

func SetupProject(prefix string) {
	http.HandleFunc(prefix+"/projects",
		middleware.CorsHeaderMiddleware(
			middleware.JsonHeaderMiddleware(
				func(w http.ResponseWriter, r *http.Request) {
					switch r.Method {
					case "GET":
						handlers.GetProjects(w, r)
					case "POST":
						handlers.PostProject(w, r)
					}
				})))
	http.Handle(prefix+"/projects/",
		http.StripPrefix(prefix+"/projects/", http.HandlerFunc(
			middleware.CorsHeaderMiddleware(
				middleware.JsonHeaderMiddleware(
					func(w http.ResponseWriter, r *http.Request) {
						// TO DO: Check if the URL.Path is valid
						switch r.Method {
						case "GET":
							handlers.GetProjectByID(w, r)
						case "DELETE":
							handlers.DeleteProject(w, r)
						}
					})))))
}

package routers

import (
	"golang-rest-api-portfolio/api/v1/handlers"
	"golang-rest-api-portfolio/api/v1/middleware"
	"net/http"
)

func SetupGithub(prefix string) {
	http.HandleFunc(prefix+"/contribution",
		middleware.CorsHeaderMiddleware(
			middleware.JsonHeaderMiddleware(
				func(w http.ResponseWriter, r *http.Request) {
					switch r.Method {
					case "GET":
						handlers.GetContribution(w, r)
					}
				})))
}

package routers

import (
	"net/http"
	"golang-rest-api-portfolio/api/v1/handlers"
)

func SetupPing(prefix string) {
	http.HandleFunc(prefix+"/ping", handlers.Ping)
}

package routers

import (
	"golang-rest-api-portfolio/api/v1/handlers"
	"net/http"
)

func SetupPing(prefix string) {
	http.HandleFunc(prefix+"/ping", handlers.Ping)
}

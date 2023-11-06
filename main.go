package main

import (
	"net/http"

	"golang-rest-api-portfolio/api/v1/routers"
)

func main() {
	routers.SetupPing("/api")
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"net/http"

	"golang-rest-api-portfolio/api/v1/routers"
)

const Port = "8080"

func main() {
	routers.SetupPing("/api")
	routers.SetupExperience("/api/v1")
	routers.SetupProject("/api/v1")
	fmt.Print("Starting http server on port " + Port)
	err := http.ListenAndServe(":"+Port, nil)
	if err != nil {
		fmt.Print(err)
	}
}

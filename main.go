package main

import (
	"flag"
	"fmt"
	"net/http"

	"golang-rest-api-portfolio/api/v1/routers"
)

const DEFAULT_PORT = "8080"

func main() {
	port := flag.String("port", DEFAULT_PORT, "Port number for our simple net/http golang Restful API")

	flag.Parse()

	routers.SetupPing("/api")
	routers.SetupExperience("/api/v1")
	routers.SetupProject("/api/v1")
	fmt.Print("Starting http server on port " + *port)
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		fmt.Print(err)
	}
}

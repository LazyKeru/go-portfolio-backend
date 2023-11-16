package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"golang-rest-api-portfolio/api/infra/service/github"
	"golang-rest-api-portfolio/api/v1/routers"
)

const DEFAULT_PORT = "8080"

func main() {
	port := flag.String("port", DEFAULT_PORT, "Port number for our simple net/http golang Restful API")
	github_api := flag.String("gapi", "test", "Github api key")

	flag.Parse()

	fmt.Print("my api" + *github_api)
	if res, err := regexp.MatchString("^ghp_[a-zA-Z0-9]{36}$", *github_api); res != true || err != nil {
		log.Fatal("Need to have a valid github api key")
	}

	github.GITHUB_API_BEARER_TOKEN = *github_api

	routers.SetupPing("/api")
	routers.SetupExperience("/api/v1")
	routers.SetupProject("/api/v1")
	routers.SetupGithub("/api/v1")
	fmt.Print("Starting http server on port " + *port)
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		fmt.Print(err)
	}
}

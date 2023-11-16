package github

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

const GITHUB_API_URL = "https://api.github.com"

var GITHUB_API_BEARER_TOKEN string

type graphqlRequest struct {
	Query string `json:"query"`
}

func GetContribution() ([]byte, error) {
	client := &http.Client{}

	query := []byte(`{
		"query": "query {viewer {name contributionsCollection {contributionCalendar {colors totalContributions weeks {contributionDays {color contributionCount date weekday}firstDay}}}}}"
	}`)

	// Need to process it to avoid the error caused by characters like \n
	// marshalledQuery, err := json.Marshal(graphqlRequest{Query: query})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Print(marshalledQuery)

	fmt.Print()

	req, err := http.NewRequest(http.MethodPost, GITHUB_API_URL+"/graphql", bytes.NewBuffer(query))
	if err != nil {
		log.Fatal(err)
	}

	// req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "bearer "+GITHUB_API_BEARER_TOKEN)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close() // Ensures that the response body will be closed after the rest of the function executes

	resbody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return resbody, err
}

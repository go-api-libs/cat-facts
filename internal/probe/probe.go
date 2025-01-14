package main

import (
	"net/http"

	"gopkg.in/dnaeon/go-vcr.v3/cassette"
)

// probe calls the API server to check what we can do
func probe() error {
	// define http calls here, e.g.: http.Get(defaultServerURL + "my-endpoint")
	_, err := http.Get("https://cat-fact.herokuapp.com/facts/591f9890d369931519ce3564")
	return err
}

// mask any secrets the API might return, e.g. in the response body
func maskSecrets(i *cassette.Interaction) error {
	return nil
}

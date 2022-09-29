package utils

import (
	"log"
	"net/http"

	"github.com/zrwaite/github-graphs/config"
)

func WakatimeGetRequest(url string) (*http.Response, error) {
	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	// add authorization header to the req
	req.Header.Add("Authorization", "Bearer "+config.CONFIG.WakatimeAccessToken)
	req.Header.Add("Host", "wakatime.com")

	// Send req using http Client
	client := &http.Client{}
	return client.Do(req)
}

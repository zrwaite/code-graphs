package utils

import (
	"log"
	"net/http"
)

func WakatimeGetRequest(url, token string) (*http.Response, error) {
	req, err := http.NewRequest("GET", "https://wakatime.com/api/v1"+url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Host", "wakatime.com")

	client := &http.Client{}
	return client.Do(req)
}

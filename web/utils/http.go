package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/zrwaite/github-graphs/config"
	"github.com/zrwaite/github-graphs/models"
	"github.com/zrwaite/github-graphs/utils/mail"
)

func WakatimeGetRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	req.Header.Add("Authorization", "Bearer "+config.CONFIG.WakatimeAccessToken)
	req.Header.Add("Host", "wakatime.com")

	client := &http.Client{}
	return client.Do(req)
}

func RefreshWakatimeToken() error {
	form := url.Values{}
	form.Add("refresh_token", config.CONFIG.WakatimeRefreshToken)
	form.Add("grant_type", "refresh_token")
	form.Add("client_id", config.CONFIG.WakatimeClientId)
	form.Add("client_secret", config.CONFIG.WakatimeClientSecret)
	form.Add("redirect_uri", config.CONFIG.RedirectURI)

	req, err := http.NewRequest("POST", "https://wakatime.com/oauth/token", strings.NewReader(form.Encode()))
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	// add authorization header to the req
	req.Header.Add("Host", "wakatime.com")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		mail.ErrorMessage("Failed to refresh wakatime token")
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("Error refreshing token: " + resp.Status)
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data))
		log.Fatal(resp)
	}
	var responseData models.WakaTimeRefreshResponse
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		mail.ErrorMessage("Failed to refresh wakatime token")
		return err
	}
	config.CONFIG.WakatimeAccessToken = responseData.AccessToken
	return nil
}

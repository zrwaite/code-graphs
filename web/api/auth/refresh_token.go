package auth

import (
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

func RefreshWakatimeToken(refreshToken string, name string) (response *models.WakaTimeTokenResponse, err error) {
	form := url.Values{}
	form.Add("refresh_token", refreshToken)
	form.Add("grant_type", "refresh_token")
	form.Add("client_id", config.CONFIG.WakatimeClientId)
	form.Add("client_secret", config.CONFIG.WakatimeClientSecret)
	form.Add("redirect_uri", config.CONFIG.AppURI+"/oauth")

	req, err := http.NewRequest("POST", "https://wakatime.com/oauth/token", strings.NewReader(form.Encode()))
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	req.Header.Add("Host", "wakatime.com")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		mail.ErrorMessage("Failed to refresh wakatime token for " + name + " - client.Do")
		return nil, err
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		fmt.Println("Error refreshing token: " + resp.Status)
		fmt.Println(string(data))
		log.Fatal(resp)
	}
	var responseData models.WakaTimeTokenResponse
	err = responseData.ParseFromString(string(data))

	if err != nil {
		mail.ErrorMessage("Failed to refresh wakatime token for " + name + " - ")
		return nil, err
	}

	return &responseData, nil
}

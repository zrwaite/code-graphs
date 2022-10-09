package auth

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

func RefreshWakatimeToken(refreshToken string) (accessToken string, err error) {
	form := url.Values{}
	form.Add("refresh_token", refreshToken)
	form.Add("grant_type", "refresh_token")
	form.Add("client_id", config.CONFIG.WakatimeClientId)
	form.Add("client_secret", config.CONFIG.WakatimeClientSecret)
	form.Add("redirect_uri", config.CONFIG.RedirectURI)

	req, err := http.NewRequest("POST", "https://wakatime.com/oauth/token", strings.NewReader(form.Encode()))
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	req.Header.Add("Host", "wakatime.com")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		mail.ErrorMessage("Failed to refresh wakatime token")
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("Error refreshing token: " + resp.Status)
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data))
		log.Fatal(resp)
	}
	var responseData models.WakaTimeTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		mail.ErrorMessage("Failed to refresh wakatime token")
		return "", err
	}

	return responseData.AccessToken, nil
}

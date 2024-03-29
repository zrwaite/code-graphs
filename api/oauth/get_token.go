package oauth

import (
	"fmt"
	"io"

	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/zrwaite/github-graphs/config"
	"github.com/zrwaite/github-graphs/models"
	"github.com/zrwaite/github-graphs/utils/mail"
)

func GetWakatimeToken(code string) (*models.WakaTimeTokenResponse, error) {
	form := url.Values{}
	form.Add("code", code)
	form.Add("grant_type", "authorization_code")
	form.Add("client_id", config.CONFIG.WakatimeClientId)
	form.Add("client_secret", config.CONFIG.WakatimeClientSecret)
	form.Add("redirect_uri", config.CONFIG.AppURI+"/oauth")

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
		fmt.Println("Failed to get wakatime token")
		return nil, err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		fmt.Println("Error getting wakatime token: " + resp.Status)
		fmt.Println(string(data))
	}
	var responseData models.WakaTimeTokenResponse
	err = responseData.ParseFromString(string(data))
	if err != nil {
		mail.ErrorMessage("Failed to get decode token")
		return nil, err
	}
	return &responseData, nil
}

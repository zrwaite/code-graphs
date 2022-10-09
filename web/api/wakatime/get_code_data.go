package wakatime

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/zrwaite/github-graphs/api/auth"
	"github.com/zrwaite/github-graphs/models"
	"github.com/zrwaite/github-graphs/utils"
	"github.com/zrwaite/github-graphs/utils/mail"
)

func getCodeData(user *models.User) (models.WakatimeData, error) {
	apiLink := "/users/current/stats/last_year?timeout=15&writes_only=true"
	resp, err := utils.WakatimeGetRequest(apiLink, user.AccessToken)
	if err != nil {
		fmt.Println(err)
	}
	var data models.WakatimeData
	if resp.StatusCode == 401 {
		fmt.Println("Refreshing token - " + time.Now().Format("2006-01-02 15:04:05"))
		accessToken, err := auth.RefreshWakatimeToken(user.RefreshToken)
		if err != nil {
			fmt.Println(err)
			return data, err
		}
		user.AccessToken = accessToken
		return getCodeData(user)
	} else if resp.StatusCode != 200 {
		fmt.Println("Error getting data: " + time.Now().Format("2006-01-02 15:04:05"))
		mail.ErrorMessage(fmt.Sprintf("Failed to get code data: \n\n\n%+v\n\n\n<img src=\"https://graphs.insomnizac.xyz/api/wakatime/pi\" />", resp))
		return data, errors.New("failed to get code data")
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	data.Expires = time.Now().Add(time.Hour * 1).Format(time.RFC1123)
	data.LastModified = time.Now().Format(time.RFC1123)
	data.Verified = true
	return data, nil
}

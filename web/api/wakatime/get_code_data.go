package wakatime

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/zrwaite/github-graphs/api/auth"
	"github.com/zrwaite/github-graphs/db/db_service"
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
		responseData, err := auth.RefreshWakatimeToken(user.RefreshToken, user.Username)
		if err != nil {
			fmt.Println(err)
			return data, err
		}
		user.AccessToken = responseData.AccessToken
		user.RefreshToken = responseData.RefreshToken
		err = db_service.UpdateUser(user)
		if err != nil {
			mail.ErrorMessage(fmt.Sprintf("Failed to update user: \n\n%+v\n", user))
			fmt.Println(err)
			return data, err
		}
		return getCodeData(user)
	} else if resp.StatusCode != 200 && resp.StatusCode != 201 && resp.StatusCode != 202 {
		fmt.Println("Error getting data: " + time.Now().Format("2006-01-02 15:04:05"))
		mail.ErrorMessage(fmt.Sprintf("Failed to get code data: \n\n\n%+v\n\n\n<img src=\"https://graphs.insomnizac.xyz/api/wakatime/pi\" />", resp))
		return data, errors.New("failed to get code data")
	}
	if resp.StatusCode == 202 {
		fmt.Println("Thicc data - " + time.Now().Format("2006-01-02 15:04:05"))
		mail.ErrorMessage("Thicc data for user " + user.Username)
		return data, errors.New("data is too thicc: graphs not loaded")
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

package oauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/zrwaite/github-graphs/models"
	"github.com/zrwaite/github-graphs/utils"
	"github.com/zrwaite/github-graphs/utils/mail"
)

func GetWakatimeUser(token string) (*models.WakaTimeUserResponse, error) {
	resp, err := utils.WakatimeGetRequest("/users/current", token)
	if err != nil {
		fmt.Println(err)
	}
	var data *models.WakaTimeUserResponse
	if resp.StatusCode != 200 {
		fmt.Println("Error getting user data: " + time.Now().Format("2006-01-02 15:04:05"))
		mail.ErrorMessage(fmt.Sprintf("Failed to get user data: \n\n%+v\n", resp))
		return data, errors.New("failed to get user data")
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

package cron

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/zrwaite/github-graphs/mail"
	"github.com/zrwaite/github-graphs/models"
	"github.com/zrwaite/github-graphs/utils"
)

func getCodeData() (models.WakatimeData, error) {
	apiLink := "https://wakatime.com/api/v1/users/current/stats/last_year?timeout=15&writes_only=true"
	resp, err := utils.WakatimeGetRequest(apiLink)
	if err != nil {
		fmt.Println(err)
	}
	var data models.WakatimeData
	if resp.StatusCode == 401 {
		fmt.Println("Refreshing token - " + time.Now().Format("2006-01-02 15:04:05"))
		err := utils.RefreshWakatimeToken()
		if err != nil {
			fmt.Println(err)
			return data, err
		}
		return getCodeData()
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
	return data, nil
}

func parseCodeData() {
	data, err := getCodeData() //wakatime token
	if err != nil {
		fmt.Println(err)
		return
	}
	writeCodeData(data)
	fmt.Println("Languages saved! Total time: " + fmt.Sprint(int(data.Data.TotalSeconds)) + " - " + time.Now().Format("2006-01-02 15:04:05"))
}

func writeCodeData(data models.WakatimeData) {
	content, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	err = utils.WriteFile("json/data.json", content)
	if err != nil {
		fmt.Println(err)
	}
}

package cron

import (
	"fmt"

	"github.com/zrwaite/github-graphs/api/wakatime"
	"github.com/zrwaite/github-graphs/db/db_service"
	"github.com/zrwaite/github-graphs/models"
)

func parseCodeData() {
	users, err := db_service.GetUsers()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, user := range users {
		if user.Verified {
			wakatime.SetCodeData(user)
		} else {
			wakatime.WriteCodeData(user.Username, models.WakatimeData{})
		}
	}
}

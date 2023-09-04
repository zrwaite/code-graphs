package wakatime

import (
	"fmt"

	"github.com/zrwaite/github-graphs/db/db_service"
	"github.com/zrwaite/github-graphs/models"
)

func ParseCodeData() {
	users, err := db_service.GetUsers()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, user := range users {
		if user.Verified {
			SetCodeData(user)
		} else {
			WriteCodeData(user.Username, models.WakatimeData{})
		}
	}
}

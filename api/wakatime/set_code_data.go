package wakatime

import (
	"fmt"
	"time"

	"github.com/zrwaite/github-graphs/db"
	"github.com/zrwaite/github-graphs/models"
)

func WriteCodeData(username string, data models.WakatimeData) error {
	err := db.SetJsonCacheNoExpire("wakatime_"+username, data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if data.Verified {
		fmt.Println("Languages saved for " + username + "! Total time: " + fmt.Sprint(int(data.Data.TotalSeconds)) + " - " + time.Now().Format("2006-01-02 15:04:05"))
	}
	return nil
}

func SetCodeData(user *models.User) error {
	data, err := getCodeData(user) //wakatime token
	if err != nil {
		fmt.Println(err)
		return err
	}
	return WriteCodeData(user.Username, data)
}

package wakatime

import (
	"github.com/zrwaite/github-graphs/db"
	"github.com/zrwaite/github-graphs/models"
)

func ReadCodeData(username string) (found bool, codeData *models.WakatimeData) {
	codeData = &models.WakatimeData{}
	found = db.GetJsonCache("wakatime_"+username, codeData)
	return found, codeData
}

package wakatime

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/zrwaite/github-graphs/models"
	"github.com/zrwaite/github-graphs/utils"
)

func ReadLanguages() models.Languages {
	languagesJson, err := utils.OpenFile("json/languages.json")
	if err != nil {
		fmt.Println("Failed to read file: " + err.Error())
	}
	defer languagesJson.Close()
	languagesBytes, _ := io.ReadAll(languagesJson)
	languages := models.Languages{}
	err = json.Unmarshal(languagesBytes, &languages)
	if err != nil {
		fmt.Println("Failed to parse json" + err.Error())
	}
	return languages
}

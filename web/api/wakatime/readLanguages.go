package wakatime

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/zrwaite/github-graphs/models"
	"github.com/zrwaite/github-graphs/utils"
)

func ReadCodeData() models.WakatimeData {
	codeDataJson, err := utils.OpenFile("json/data.json")
	if err != nil {
		fmt.Println("Failed to read file: " + err.Error())
	}
	defer codeDataJson.Close()
	codeDataBytes, _ := io.ReadAll(codeDataJson)
	codeData := models.WakatimeData{}
	err = json.Unmarshal(codeDataBytes, &codeData)
	if err != nil {
		fmt.Println("Failed to parse json" + err.Error())
	}
	return codeData
}

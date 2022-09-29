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
		fmt.Println("Invalid token")
		mail.ErrorMessage("Invalid token on code graphs")
	} else if resp.StatusCode != 200 {
		fmt.Println("Error getting data: " + time.Now().Format("2006-01-02 15:04:05"))
		mail.ErrorMessage(fmt.Sprintf("Failed to get code data: \n\n\n%+v\n\n\n<img src=\"https://graphs.insomnizac.xyz/api/wakatime/pi\" />", resp))
		return data, errors.New("failed to get code data")
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func parseLanguages() {
	ignoreLanguages := []string{"JSON", "Docker", "Markdown", "Other", "INI", "Text", "XML", "YAML", "Bash", "Git Config", "Objective-C", "TOML", "Apache Config", "GitIgnore file", "Shell Script", "GraphQL"}
	data, err := getCodeData() //wakatime token
	if err != nil {
		fmt.Println(err)
		return
	}
	languages := []models.Language{}
	other := models.Language{
		Name:         "Other",
		TotalSeconds: 0,
		Colour:       "white",
		Percent:      0,
	}
	totalPercent := 100.0
	totalSeconds := data.Data.TotalSeconds
	for _, language := range data.Data.Languages {
		if utils.Contains(ignoreLanguages, language.Name) {
			totalPercent -= language.Percent
			totalSeconds -= language.TotalSeconds
			continue
		}
		if len(languages) > 11 {
			other.Percent += language.Percent
			other.TotalSeconds += language.TotalSeconds
			continue
		}
		languages = append(languages, models.Language{
			Name:         language.Name,
			Colour:       getColour(language.Name),
			TotalSeconds: language.TotalSeconds,
			Percent:      language.Percent,
		})
	}
	if other.TotalSeconds > 0 {
		languages = append(languages, other)
	}
	for i := 0; i < len(languages); i++ {
		languages[i].Percent = languages[i].Percent / totalPercent * 100
	}
	writeLanguages(models.Languages{
		Languages:    languages,
		TotalSeconds: data.Data.TotalSeconds,
	})
	fmt.Println("Languages saved! Total time: " + fmt.Sprint(int(data.Data.TotalSeconds)) + " - " + time.Now().Format("2006-01-02 15:04:05"))
}

func writeLanguages(data models.Languages) {
	content, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	err = utils.WriteFile("json/languages.json", content)
	if err != nil {
		fmt.Println(err)
	}
}

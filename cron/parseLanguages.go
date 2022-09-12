package cron

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/zrwaite/github-graphs/models"
	"github.com/zrwaite/github-graphs/utils"
)

func getCodeData() (models.WakatimeData, error) {
	apiLink := "https://wakatime.com/api/v1/users/current/summaries?timeout=15&writes_only=true"
	// startDate := time.Now().AddDate(-1, 0, 0).Format("2006-01-02")
	startDate := "2021/10/12"
	date := time.Now().Format("2006-01-02")
	apiLink += "&start=" + startDate + "&end=" + date
	resp, err := utils.WakatimeGetRequest(apiLink)
	if err != nil {
		fmt.Println(err)
	}
	var data models.WakatimeData

	if resp.StatusCode == 401 {
		fmt.Println("Invalid token")
	} else if resp.StatusCode != 200 {
		log.Fatal("Error getting data")
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
	totalTime := 0
	foundLanguages := []models.Language{}
	for _, datapoint := range data.Data {
		for _, language := range datapoint.Languages {
			found := false
			time := int(language.TotalSeconds)
			for i := 0; i < len(foundLanguages); i++ {
				if language.Name != foundLanguages[i].Name {
					continue
				}
				foundLanguages[i].Time += time
				totalTime += time
				found = true
				break
			}
			if !found {
				if !utils.Contains(ignoreLanguages, language.Name) {
					foundLanguages = append(foundLanguages, models.Language{
						Name:   language.Name,
						Time:   time,
						Colour: getColour(language.Name),
					})
				}
			}
		}
	}
	languages := utils.SortLanguagesByTime(foundLanguages)
	if len(languages) > 13 {
		// for i := 13; i < len(languages); i++ {
		// 	totalTime -= languages[i].Time
		// }
		languages = languages[:13]
	}
	for i := 0; i < len(languages); i++ {
		languages[i].Percent = 100 * float64(languages[i].Time) / float64(totalTime)
	}
	writeLanguages(models.Languages{
		Languages: languages,
		TotalTime: totalTime,
	})
	fmt.Println("Languages saved: " + time.Now().Format("2006-01-02 15:04:05"))
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

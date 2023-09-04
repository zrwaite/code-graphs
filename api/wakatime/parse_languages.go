package wakatime

import (
	"github.com/zrwaite/github-graphs/models"
	"github.com/zrwaite/github-graphs/utils"
)

func ParseLanguages(data *models.WakatimeData, ignoreLanguages []string) models.Languages {
	languageData := models.Languages{}
	totalPercent := 100.0
	totalSeconds := data.Data.TotalSeconds
	for _, language := range data.Data.Languages {
		if utils.UncapitalizedContains(ignoreLanguages, language.Name) {
			totalPercent -= language.Percent
			totalSeconds -= language.TotalSeconds
			continue
		}
		if language.Name == "Other" || len(languageData.Languages) > 11 {
			languageData.Other.Languages = append(languageData.Other.Languages, models.Language{
				Name:         language.Name,
				Colour:       utils.GetColour(language.Name),
				TotalSeconds: language.TotalSeconds,
				Percent:      language.Percent,
			})
			continue
		}
		languageData.Languages = append(languageData.Languages, models.Language{
			Name:         language.Name,
			Colour:       utils.GetColour(language.Name),
			TotalSeconds: language.TotalSeconds,
			Percent:      language.Percent,
		})
	}
	for i := 0; i < len(languageData.Languages); i++ {
		languageData.Languages[i].Percent = languageData.Languages[i].Percent / totalPercent * 100
	}
	for i := 0; i < len(languageData.Other.Languages); i++ {
		languageData.Other.Percent += languageData.Other.Languages[i].Percent
		languageData.Other.TotalSeconds += languageData.Other.Languages[i].TotalSeconds
		languageData.Other.Languages[i].Percent = languageData.Other.Languages[i].Percent / totalPercent * 100
	}
	languageData.Other.Percent = languageData.Other.Percent / totalPercent * 100

	return languageData
}

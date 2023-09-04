package wakatime_top_stats

import (
	"github.com/zrwaite/github-graphs/models"
	"github.com/zrwaite/github-graphs/utils"
)

func ParseStatsData(stats models.WakatimeStatsData, ignoreList []string) (languages []string, editors []string, operatingSystems []string, projects []string) {
	for _, language := range stats.Languages {
		if !utils.UncapitalizedContains(ignoreList, language.Name) {
			languages = append(languages, language.Name)
		}
		if len(languages) >= 3 {
			break
		}
	}
	for _, editor := range stats.Editors {
		if !utils.UncapitalizedContains(ignoreList, editor.Name) {
			editors = append(editors, editor.Name)
		}
		if len(editors) >= 3 {
			break
		}
	}
	for _, operatingSystem := range stats.OperatingSystems {
		if !utils.UncapitalizedContains(ignoreList, operatingSystem.Name) {
			operatingSystems = append(operatingSystems, operatingSystem.Name)
		}
		if len(operatingSystems) >= 3 {
			break
		}
	}
	for _, project := range stats.Projects {
		if !utils.UncapitalizedContains(ignoreList, project.Name) {
			projects = append(projects, project.Name)
		}
		if len(projects) >= 3 {
			break
		}
	}
	return
}

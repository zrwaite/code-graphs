package wakatime

import (
	"net/http"
	"strings"

	svg "github.com/ajstarks/svgo"
)

func GetWakatimePiSVG(w http.ResponseWriter, r *http.Request) {
	codeData := ReadCodeData()
	w.Header().Set("Last-Modified", codeData.LastModified)
	w.Header().Set("Expires", codeData.Expires)
	ignoreString := r.URL.Query().Get("ignore")
	addDefaultIgnore := r.URL.Query().Get("addDefaultIgnore")
	ignoreLanguages := []string{}
	if ignoreString != "" {
		ignoreLanguages = strings.Split(ignoreString, ",")
	}
	if addDefaultIgnore == "true" {
		defaultIgnoreLanguages := []string{"JSON", "Markdown", "Other", "INI", "Text", "XML", "YAML", "Bash", "Git Config", "TOML", "Apache Config", "GitIgnore file", "Shell Script", "GraphQL", "Tex", "CMake", "Git"}
		ignoreLanguages = append(ignoreLanguages, defaultIgnoreLanguages...)
	}
	languages := parseLanguages(codeData, ignoreLanguages)
	s := svg.New(w)
	s.Start(1500, 917)
	s.Roundrect(0, 0, 1500, 917, 20, 20, "fill:black;stroke:white;stroke-width:6")
	CreatePiGraph(s, languages)
	CreateLegend(s, languages)
	s.End()
}

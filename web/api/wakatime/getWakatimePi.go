package wakatime

import (
	"strings"

	svg "github.com/ajstarks/svgo"
	"github.com/gin-gonic/gin"
)

func GetWakatimePiSVG(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.String(405, "Method not allowed")
		return
	}
	codeData := ReadCodeData()
	c.Header("Last-Modified", codeData.LastModified)
	c.Header("Expires", codeData.Expires)

	ignoreString := c.Query("ignore")
	removeDefaultIgnore := c.Query("removeDefaultIgnore")

	ignoreLanguages := []string{}
	if ignoreString != "" {
		ignoreLanguages = strings.Split(ignoreString, ",")
	}
	if removeDefaultIgnore != "true" {
		defaultIgnoreLanguages := []string{"JSON", "Markdown", "Other", "INI", "Text", "XML", "YAML", "Git Config", "TOML", "Apache Config", "GitIgnore file", "GraphQL", "Tex", "CMake", "Git"}
		ignoreLanguages = append(ignoreLanguages, defaultIgnoreLanguages...)
	}
	languages := parseLanguages(codeData, ignoreLanguages)

	c.Header("Content-Type", "image/svg+xml")
	c.Header("Cache-Control", "public, max-age=3600")
	s := svg.New(c.Writer)
	s.Start(1500, 917)
	s.Roundrect(0, 0, 1500, 917, 20, 20, "fill:black;stroke:white;stroke-width:6")
	CreatePiGraph(s, languages)
	CreateLegend(s, languages)
	s.End()
}

package wakatime_pi

import (
	"strings"

	svg "github.com/ajstarks/svgo"
	"github.com/gin-gonic/gin"
	"github.com/zrwaite/github-graphs/api/wakatime"
)

func GetWakatimePiSVG(c *gin.Context) {
	username := c.Params.ByName("username")
	found, codeData := wakatime.ReadCodeData(username)
	if !codeData.Verified {
		c.String(401, "User not verified")
		return
	}
	if !found {
		c.String(401, "User not found")
		return
	}
	c.Header("Last-Modified", codeData.LastModified)
	c.Header("Expires", codeData.Expires)

	ignoreString := c.Query("ignore")
	removeDefaultIgnore := c.Query("removeDefaultIgnore")
	addUsername := c.Query("addUsername") == "true"

	ignoreLanguages := []string{}
	if ignoreString != "" {
		ignoreLanguages = strings.Split(ignoreString, ",")
	}
	if removeDefaultIgnore != "true" {
		defaultIgnoreLanguages := []string{"JSON", "Markdown", "Other", "INI", "Text", "XML", "YAML", "Git Config", "TOML", "Apache Config", "GitIgnore file", "GraphQL", "Tex", "CMake", "Git", "George", "CSV", "textmate", "roff", "Tcsh", "sbt", "Nginx configuration", "Protocol Buffer", "Properties"}
		ignoreLanguages = append(ignoreLanguages, defaultIgnoreLanguages...)
	}
	languages := wakatime.ParseLanguages(codeData, ignoreLanguages)

	c.Header("Content-Type", "image/svg+xml")
	c.Header("Cache-Control", "public, max-age=3600")
	s := svg.New(c.Writer)
	s.Start(1500, 917)
	s.Roundrect(0, 0, 1500, 917, 20, 20, "fill:black;stroke:white;stroke-width:6")
	CreatePiGraph(s, languages, addUsername, username)
	CreateLegend(s, languages)
	s.End()
}

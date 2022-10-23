package wakatime_top_stats

import (
	"strings"

	svg "github.com/ajstarks/svgo"
	"github.com/gin-gonic/gin"
	"github.com/zrwaite/github-graphs/api/wakatime"
)

func GetWakatimeTopStatsSVG(c *gin.Context) {
	username := c.Params.ByName("username")
	found, codeData := wakatime.ReadCodeData(username)
	if !found {
		c.String(401, "User not found")
		return
	}
	if !codeData.Verified {
		c.String(401, "User not verified")
		return
	}
	c.Header("Last-Modified", codeData.LastModified)
	c.Header("Expires", codeData.Expires)

	ignoreString := c.Query("ignore")
	removeDefaultIgnore := c.Query("removeDefaultIgnore")
	addUsername := c.Query("addUsername") == "true"

	ignoreList := []string{}
	if ignoreString != "" {
		ignoreList = strings.Split(ignoreString, ",")
	}
	if removeDefaultIgnore != "true" {
		defaultIgnoreLanguages := []string{"JSON", "Markdown", "Other", "INI", "Text", "XML", "YAML", "Git Config", "TOML", "Apache Config", "GitIgnore file", "GraphQL", "Tex", "CMake", "Git", "George", "CSV", "textmate", "roff", "Tcsh", "sbt", "Nginx configuration", "Protocol Buffer", "Properties"}
		ignoreList = append(ignoreList, defaultIgnoreLanguages...)
	}

	languages, editors, operatingSystems, projects := ParseStatsData(codeData.Data, ignoreList)

	c.Header("Content-Type", "image/svg+xml")
	c.Header("Cache-Control", "public, max-age=3600")
	s := svg.New(c.Writer)
	s.Start(1500, 917)
	s.Roundrect(0, 0, 1500, 917, 20, 20, "fill:black;stroke:white;stroke-width:6")
	y_offset := 0
	if addUsername {
		y_offset = 50
		s.Text(750, 50, username+"'s Top Wakatime Stats", "text-anchor:middle;fill:white;font-size:40px;font-family:monospace")
	}
	CreateList(s, "Top Languages", languages, 375, 50, 80+y_offset, true)
	CreateList(s, "Top Editors", editors, 1125, 800, 80+y_offset, true)
	CreateList(s, "Top OS's", operatingSystems, 375, 50, 400+y_offset, false)
	CreateList(s, "Top Projects", projects, 1125, 800, 400+y_offset, false)
	s.End()
}

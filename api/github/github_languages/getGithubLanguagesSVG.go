package github_languages

import (
	"github.com/gin-gonic/gin"
)

func GetGithubLanguagesSVG(c *gin.Context) {
	username := c.Params.ByName("username")
	hide := c.Query("hide")

	if username == "" {
		c.String(400, "Username is required")
		return
	}

	html, err := GetGithubLanguagesData(username, hide)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	svg := GetSvg(html)
	graphBytes := []byte(svg)

	c.Header("Content-Type", "image/svg+xml")
	c.Header("Cache-Control", "public, max-age=3600")
	c.Data(200, "image/svg+xml", graphBytes)
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zrwaite/github-graphs/config"
)

func HomeController(c *gin.Context) {
	prefix := config.CONFIG.AppURI
	username := c.Params.ByName("username")
	wakatime_pi_graph_link := prefix + "/api/wakatime/Insomnizac"
	var authorized bool
	if username != "" {
		wakatime_pi_graph_link = prefix + "/api/wakatime/" + username
		authorized = true
	}
	c.HTML(http.StatusOK, "index.go.tmpl", gin.H{
		"github_streak_graph_link": prefix + "/api/streak/zrwaite",
		"wakatime_pi_graph_link":   wakatime_pi_graph_link,
		"authorized":               authorized,
		"auth_href":                "https://wakatime.com/oauth/authorize?client_id=YrKmoBVz3M8lDJZmTFaFWDKz&response_type=code&scope=read_stats,read_logged_time,email&redirect_uri=" + prefix + "/oauth",
	})
}

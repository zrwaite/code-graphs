package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zrwaite/github-graphs/config"
	"github.com/zrwaite/github-graphs/db/db_service"
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

	users, err := db_service.GetUsersWithCache()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	public_graph_links := []string{}
	for i := len(users) - 1; i >= 0; i-- {
		if users[i].Public {
			public_graph_links = append(public_graph_links, prefix+"/api/wakatime/"+users[i].Username+"?addUsername=true")
		}
	}

	c.HTML(http.StatusOK, "index.go.tmpl", gin.H{
		"github_streak_graph_link":    prefix + "/api/streak/zrwaite",
		"github_languages_graph_link": prefix + "/api/github/languages/zrwaite?hide=html,css",
		"wakatime_pi_graph_link":      wakatime_pi_graph_link,
		"authorized":                  authorized,
		"auth_href":                   "https://wakatime.com/oauth/authorize?client_id=YrKmoBVz3M8lDJZmTFaFWDKz&response_type=code&scope=read_stats,read_logged_time,email&redirect_uri=" + prefix + "/oauth",
		"public_graph_links":          public_graph_links,
	})
}

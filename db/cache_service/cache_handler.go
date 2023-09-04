package cache_service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zrwaite/github-graphs/api/wakatime"
	"github.com/zrwaite/github-graphs/config"
	"github.com/zrwaite/github-graphs/db"
)

func ClearCacheHandler(c *gin.Context) {
	var json struct {
		Password string `json:"password" binding:"required"`
	}

	if c.Bind(&json) == nil {
		if json.Password == config.CONFIG.AdminPassword {
			db.ClearCache()
			wakatime.ParseCodeData()
			c.JSON(200, gin.H{"status": "ok"})
		} else {
			c.JSON(401, gin.H{"status": "Invalid password"})
		}
	} else {
		c.String(http.StatusBadRequest, "Must provide password")
	}
}

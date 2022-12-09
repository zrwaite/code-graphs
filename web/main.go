package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/zrwaite/github-graphs/api/oauth"
	"github.com/zrwaite/github-graphs/api/streak"
	"github.com/zrwaite/github-graphs/api/wakatime/wakatime_pi"
	"github.com/zrwaite/github-graphs/api/wakatime/wakatime_top_stats"
	"github.com/zrwaite/github-graphs/client/controllers"
	"github.com/zrwaite/github-graphs/config"
	"github.com/zrwaite/github-graphs/cron"
	"github.com/zrwaite/github-graphs/db"
	"github.com/zrwaite/github-graphs/db/cache_service"
)

// var db = make(map[string]string)

const port = "8001"

func main() {
	r := setupRouter()
	fmt.Println("Starting server at http://localhost:" + port)
	r.Run(":" + port)
}

func setupRouter() *gin.Engine {
	godotenv.Load(".env")
	config.ConfigInit()
	db.ConnectToMongoDB()
	db.ConnectToRedis()
	// db.InitializeDatabase()
	go cron.RunCronJobs()
	// mail.StartupMessage()

	r := gin.Default()
	r.LoadHTMLGlob("client/templates/*/*.go.tmpl")

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/", controllers.HomeController)
	r.GET("/:username", controllers.HomeController)
	r.Static("/styles", "./client/static/css")
	r.Static("/fonts", "./client/static/fonts")
	r.GET("/api/streak/:username", streak.GetStreakSVG)
	r.GET("/api/wakatime/:username", wakatime_pi.GetWakatimePiSVG)
	r.GET("/api/wakatime_stats/:username", wakatime_top_stats.GetWakatimeTopStatsSVG)
	r.GET("/oauth", oauth.OAuthHandler)
	r.POST("/clear_cache", cache_service.ClearCacheHandler)

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	// authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
	// 	"foo":  "bar", // user:foo password:bar
	// 	"manu": "123", // user:manu password:123
	// }))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	// authorized.POST("admin", func(c *gin.Context) {
	// 	user := c.MustGet(gin.AuthUserKey).(string)

	// 	// Parse JSON
	// 	var json struct {
	// 		Value string `json:"value" binding:"required"`
	// 	}

	// 	if c.Bind(&json) == nil {
	// 		db[user] = json.Value
	// 		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	// 	}
	// })

	return r
}

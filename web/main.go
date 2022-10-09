package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/zrwaite/github-graphs/api"
	"github.com/zrwaite/github-graphs/client/controllers"
	"github.com/zrwaite/github-graphs/config"
)

// var db = make(map[string]string)

const port = "8001"

func main() {
	r := setupRouter()
	r.Run(":" + port)
}

func setupRouter() *gin.Engine {
	godotenv.Load(".env")
	config.ConfigInit()
	fmt.Println("Starting server at http://localhost:" + port)
	// go cron.RunCronJobs()
	// mail.StartupMessage()

	r := gin.Default()
	r.LoadHTMLGlob("client/templates/*/*.go.tmpl")

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/", controllers.HomeController)
	// r.GET("/styles", http.FileServer(http.Dir("../client")))
	r.Static("/styles", "./static/css")
	r.GET("/api/:type", api.NewAPIHandler)

	// Get user value
	// r.GET("/user/:name", func(c *gin.Context) {
	// 	user := c.Params.ByName("name")
	// 	value, ok := db[user]
	// 	if ok {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	// 	} else {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	// 	}
	// })

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

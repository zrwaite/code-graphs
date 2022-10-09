package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeController(c *gin.Context) {
	c.HTML(http.StatusOK, "index.go.tmpl", gin.H{
		"title": "Code Graphs",
	})
}

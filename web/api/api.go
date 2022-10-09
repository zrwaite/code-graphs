package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zrwaite/github-graphs/api/streak"
	"github.com/zrwaite/github-graphs/api/wakatime"
)

func NewAPIHandler(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.String(http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	switch c.Request.URL.Path {
	case "/api/streak":
		streak.GetStreakSVG(c)
	case "/api/wakatime":
		wakatime.GetWakatimePiSVG(c)
	case "/api/test":
		c.String(http.StatusOK, "test")
	default:
		c.String(http.StatusNotFound, "Not found")
	}
}

// func APIHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-Type", "image/svg+xml")
// 	w.Header().Set("Cache-Control", "public, max-age=3600")
// 	if r.Method != "GET" {
// 		w.WriteHeader(400)
// 		json.NewEncoder(w).Encode("Method " + r.Method + " is not supported")
// 		return
// 	} else {
// 		switch r.URL.Path {
// 		case "/api/streak":
// 			streak.GetStreakSVG(w, r)
// 		case "/api/wakatime/pi":
// 			wakatime.GetWakatimePiSVG(w, r)
// 		default:
// 			w.WriteHeader(400)
// 			json.NewEncoder(w).Encode("Endpoint not found")
// 			return
// 		}
// 	}
// }

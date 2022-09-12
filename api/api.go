package api

import (
	"encoding/json"
	"net/http"

	"github.com/zrwaite/github-graphs/api/streak"
	"github.com/zrwaite/github-graphs/api/wakatime"
)

func APIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "image/svg+xml")
	if r.Method != "GET" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Method " + r.Method + " is not supported")
		return
	} else {
		switch r.URL.Path {
		case "/api/streak":
			streak.GetStreakSVG(w, r)
		case "/api/wakatime/pi":
			wakatime.GetWakatimePiSVG(w, r)
		default:
			w.WriteHeader(400)
			json.NewEncoder(w).Encode("Endpoint not found")
			return
		}
	}
}

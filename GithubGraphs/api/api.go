package api

import (
	"encoding/json"
	"net/http"

	"github.com/zrwaite/github-graphs/api/streak"
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
			username := r.URL.Query().Get("username")
			if username == "" {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode("username parameter is required")
				return
			}
			streakGraph, err := streak.CreateStreakGraph(username)
			if err != nil {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(err.Error())
				return
			} else {
				w.WriteHeader(200)
				streakGraphBytes := []byte(streakGraph)
				w.Write(streakGraphBytes)
				return
			}
		default:
			w.WriteHeader(400)
			json.NewEncoder(w).Encode("Endpoint not found")
			return
		}
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/zrwaite/github-graphs/api"
	"github.com/zrwaite/github-graphs/config"
	"github.com/zrwaite/github-graphs/cron"
	"github.com/zrwaite/github-graphs/mail"
)

const port = "8001"

func main() {
	godotenv.Load(".env")
	config.ConfigInit()
	http.HandleFunc("/api/", api.APIHandler)
	fmt.Println("Starting server at http://localhost:" + port)
	go cron.RunCronJobs()
	mail.StartupMessage()
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

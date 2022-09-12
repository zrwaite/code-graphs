package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zrwaite/github-graphs/api"
	"github.com/zrwaite/github-graphs/config"
)

const port = "8009"

func main() {
	config.ConfigInit()
	http.HandleFunc("/api/", api.APIHandler)
	fmt.Println("Starting server at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

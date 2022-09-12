package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zrwaite/github-graphs/api"
)

const port = "8009"

func main() {
	http.HandleFunc("/api/", api.APIHandler)
	fmt.Println("Starting server at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

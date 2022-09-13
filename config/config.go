package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var CONFIG Config

func ConfigInit() {
	initDirectory()
	initEnv()
}

func initDirectory() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	CONFIG.Directory = filepath.Dir(ex) + "/"
	if strings.Contains(ex, "var/folders") {
		fmt.Println("dev mode directory")
		CONFIG.Directory = "/Users/zacharywaite/Coding/CodeGraphs/"
	}
}

func initEnv() {
	CONFIG.WakatimeToken = os.Getenv("WAKATIME_TOKEN")
	if CONFIG.WakatimeToken == "" {
		log.Fatal("Failed to load env variables")
	}
	CONFIG.SendgridAPIKey = os.Getenv("SENDGRID_API_KEY")
	CONFIG.ContactEmail = os.Getenv("CONTACT_EMAIL")
	CONFIG.FromEmail = os.Getenv("FROM_EMAIL")
}

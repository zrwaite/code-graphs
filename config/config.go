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
	CONFIG.WakatimeAccessToken = os.Getenv("WAKATIME_ACCESS_TOKEN")
	if CONFIG.WakatimeAccessToken == "" {
		log.Fatal("Failed to load env variables")
	}
	CONFIG.WakatimeRefreshToken = os.Getenv("WAKATIME_REFRESH_TOKEN")
	CONFIG.RedirectURI = os.Getenv("REDIRECT_URI")
	CONFIG.WakatimeClientId = os.Getenv("WAKATIME_CLIENT_ID")
	CONFIG.WakatimeClientSecret = os.Getenv("WAKATIME_CLIENT_SECRET")
	CONFIG.SendGridAPIKey = os.Getenv("SENDGRID_API_KEY")
	CONFIG.ContactEmail = os.Getenv("CONTACT_EMAIL")
	CONFIG.FromEmail = os.Getenv("FROM_EMAIL")
}

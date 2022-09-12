package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var CONFIG Config

func ConfigInit() {
	initDirectory()
}

func initDirectory() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	CONFIG.Directory = filepath.Dir(ex)
	if strings.Contains(ex, "var/folders") {
		fmt.Println("dev mode directory")
		CONFIG.Directory = "/Users/zacharywaite/Coding/CodeGraphs/GoGraphs"
	}
}

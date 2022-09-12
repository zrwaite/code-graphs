package utils

import (
	"os"

	"github.com/zrwaite/github-graphs/config"
)

func WriteImage(filename string, data []byte) error {
	filePath := config.CONFIG.Directory + "/images/" + filename
	return os.WriteFile(filePath, data, 0644)
}

func OpenImage(filename string) (*os.File, error) {
	filePath := config.CONFIG.Directory + "/images/" + filename
	return os.Open(filePath)
}

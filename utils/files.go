package utils

import (
	"os"

	"github.com/zrwaite/github-graphs/config"
)

func OpenFile(filename string) (*os.File, error) {
	filePath := config.CONFIG.Directory + filename
	return os.Open(filePath)
}

func WriteFile(filename string, data []byte) error {
	filePath := config.CONFIG.Directory + filename
	return os.WriteFile(filePath, data, 0644)
}

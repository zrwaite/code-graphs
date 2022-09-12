package utils

import (
	"os"

	"github.com/zrwaite/github-graphs/config"
)

func OpenFile(filename string) (*os.File, error) {
	filePath := config.CONFIG.Directory + filename
	return os.Open(filePath)
}

package utils

import (
	"os"
)

func OpenImage(filename string) (*os.File, error) {
	return OpenFile("/images/" + filename)
}

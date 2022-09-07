package utils

import (
	"os"
	"strings"
)

func ReadInputFile(file string) (string) {
	raw, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(raw))
}

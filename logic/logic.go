package logic

import (
	"strings"
	"path/filepath"
)

const index = "index.html"

func ParsePath(path string) string  {
	if path == "" {
		return "index.html"
	}
	parsedPath := strings.Split(filepath.Clean(path), "/")
	correctPath := ""
	for _, v := range(parsedPath) {
        if len(v) > 0 {
            correctPath += v + "/"
        }
    }

	if strings.Contains(correctPath, ".") {
		correctPath = correctPath[:len(correctPath) - 1]
	} else {
		correctPath += "index.html"
	}

	return correctPath
}

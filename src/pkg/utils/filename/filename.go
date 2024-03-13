package filename

import (
	"regexp"
	"strings"
)

func GetFileName(url string) string {
	lastSlashIndex := strings.LastIndex(url, "/")
	if lastSlashIndex == -1 {
		lastSlashIndex = 0
	} else {
		lastSlashIndex++
	}

	fileName := url[lastSlashIndex:]

	queryIndex := strings.Index(fileName, "?")
	if queryIndex != -1 {
		fileName = fileName[:queryIndex]
	}

	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	fileName = re.ReplaceAllString(fileName, ".")

	return fileName + ".html"
}

package helpers

import (
	"os"
	"strings"
)

// EnforceHTTP makes every url a http
func EnforceHTTP(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

// RemoveDomainError removes all the commonly found prefixes from URL
// such as http, https, www
// then checks of the remaining string is the DOMAIN itself
func RemoveDomainError(url string) bool {
	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]

	return newURL == os.Getenv("DOMAIN")
}

package app

import (
	"net/url"
	"strings"
)

// ParseLink parses a link and returns host without www.
// It ignores query params and fragments.
// It returns string that will be the name of the saved file.
func ParseLink(link string) string {
	u, err := url.Parse(link)
	if err != nil {
		return ""
	}

	hostname := u.Hostname()
	trimmed := strings.TrimPrefix(hostname, "www.")

	return trimmed
}

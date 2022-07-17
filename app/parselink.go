package app

import (
	"strings"
)

// punctuation characters
var punctuation = []string{
	"!", "?", ",", ":", ";", "\"",
	"'", "`", "~", "@", "#", "$",
	"%", "^", "&", "*", "(", ")",
	"+", "=", "{", "}", "[", "]",
	"|", "\\", "<", ">", "/", " ",
}

// ParseLink parses the link, cuts the domain and returns first 5 characters of the last part of the link.
// It returns string that will be the name of the saved file.
func ParseLink(link string) string {
	splittedLink := strings.Split(link, "/")
	domainName := splittedLink[2]
	sectionName := splittedLink[len(splittedLink)-2]
	articleName := splittedLink[len(splittedLink)-1]
	fileName := domainName
	if articleName != "" {
		for _, char := range punctuation {
			articleName = strings.Replace(articleName, char, "", -1)
		}
		if strings.Contains(articleName, ".") {
			articleNameSplitted := strings.Split(articleName, ".")
			articleName = articleNameSplitted[0]
		}
		if len(articleName) > 10 {
			fileName += "-" + articleName[0:10]
		} else {
			fileName += "-" + articleName
		}
	} else {
		fileName += "-" + sectionName
	}
	return fileName
}

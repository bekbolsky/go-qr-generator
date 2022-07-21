package app

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// ReadLinkFromFile reads the link from the file line by line.
// It returns the links as a slice of strings
func ReadLinkFromFile(filePath string) []string {
	links := make([]string, 0)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("open file failed: %v", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line != "" {
			links = append(links, line)
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file failed: %v", err)
	}
	return links
}

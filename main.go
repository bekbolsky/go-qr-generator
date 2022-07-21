package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/bekbolsky/go-qr-generator/app"
)

func main() {

	var (
		inputFileName string
		outputFolder  string
		FgColor       string
		transparent   bool
	)

	// create flags
	flag.StringVar(&inputFileName, "input", "links.txt", "input file name")
	flag.StringVar(&outputFolder, "output", "QRcodes", "output folder")
	flag.StringVar(&FgColor, "fgcolor", "#000000", "changing foreground color")
	flag.BoolVar(&transparent, "transparent", false, "enabling transparent background")
	flag.Parse()

	// check flags
	// print help if no flags
	if inputFileName == "" || outputFolder == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// create folder if not exists
	if _, err := os.Stat(outputFolder); os.IsNotExist(err) {
		os.Mkdir(outputFolder, 0755)
	}

	// read links from file
	links := app.ReadLinkFromFile(inputFileName)

	var wg sync.WaitGroup

	// make qr codes
	for i, link := range links {
		wg.Add(1)
		go func(i int, link string) {
			defer wg.Done()

			fileName := app.ParseLink(link)
			// output in current directory
			outputFileName := fmt.Sprintf("%s/%s-%d.png", outputFolder, fileName, i+1)
			app.GenerateQR(link, outputFileName, transparent, FgColor)
		}(i, link)
	}
	wg.Wait()

	fmt.Println("QR codes created")
}

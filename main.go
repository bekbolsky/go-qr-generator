package main

import (
	"flag"
	"fmt"
	"os"

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

	// make qr codes
	for i, link := range links {
		fileName := app.ParseLink(link)
		// output in current directory
		outputFileName := fmt.Sprintf("%s/%s-%d.png", outputFolder, fileName, i+1)
		app.GenerateQR(link, outputFileName, transparent, FgColor)
	}
	fmt.Println("QR codes created")
}

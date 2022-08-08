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
		// add goroutine to wait group
		wg.Add(1)
		go func(i int, link string) {
			fileName := app.ParseLink(link)
			// output in current directory
			outputFileName := fmt.Sprintf("%s/%s-%d.png", outputFolder, fileName, i+1)
			app.GenerateQR(link, outputFileName, transparent, FgColor)
			// inform wait group that goroutine is done
			wg.Done()
		}(i, link)
	}

	// wait for all qr codes to be generated
	// and then exit
	// this is to prevent program from exiting before all qr codes are generated
	wg.Wait()

	fmt.Println("QR codes created")
}

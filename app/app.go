package app

import (
	"image/color"
	"log"

	qrcode "github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func GenerateQR(link string, output string, transparent bool, FgColor string) {
	encodeOptions := []qrcode.EncodeOption{
		qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionMedium),
		qrcode.WithEncodingMode(qrcode.EncModeAuto),
	}

	qr, err := qrcode.NewWith(link, encodeOptions...)
	if err != nil {
		log.Fatalf("initializing qrcode failed: %v", err)
	}

	var transparency standard.ImageOption
	if transparent {
		transparency = standard.WithBgColor(color.RGBA{255, 255, 255, 0})
	} else {
		transparency = standard.WithBgColor(color.White)
	}
	imageOptions := []standard.ImageOption{
		transparency,
		standard.WithBorderWidth(0),
		standard.WithFgColorRGBHex(FgColor),
		standard.WithBuiltinImageEncoder(standard.PNG_FORMAT),
	}

	w, err := standard.New(output, imageOptions...)
	if err != nil {
		log.Fatalf("creating a standard writer failed: %v", err)
	}
	if err := qr.Save(w); err != nil {
		log.Fatalf("saving qr code failed: %v", err)
	}
}

/* Simple command line program using to crop images

Usage:

$ go run crop.go --height=5000 --width=7000 <path to>/cat1.jpg  <path to>cat2.png

The cropped images will be placed in the same directory as the original images with
the file names being ``cropped_<original_file_name>.<original_extension>``
*/
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/oliamb/cutter"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func validateImageType(image_type string) bool {
	recognizedImageTypes := []string{"image/jpeg", "image/png"}
	for _, t := range recognizedImageTypes {
		if image_type == t {
			return true
		}
	}
	return false
}

/*
cropper() takes in four parameters

imageData : byte[] representing the input image

cWidth    : An int specifying the desired width of the cropped image

cHeight   : An int specifying the desired height of the cropped image
*/
func cropper(inputFilePath string, cWidth int, cHeight int) {
	imageData, _ := ioutil.ReadFile(inputFilePath)
	imageType := http.DetectContentType(imageData)

	// Check if we can handle this image
	if !validateImageType(imageType) {
		panic("Cannot handle image of this type")
	}

	// We first decode the image
	reader := bytes.NewReader(imageData)
	img, _, err := image.Decode(reader)
	if err != nil {
		panic(err)
	}

	// Perform the cropping
	croppedImg, err := cutter.Crop(img, cutter.Config{
		Height:  cHeight,
		Width:   cWidth,
		Mode:    cutter.TopLeft,
		Anchor:  image.Point{60, 10},
		Options: 0,
	})

	if err != nil {
		panic(err)
	}

	// Now we encode the cropped image data using the appropriate
	// encoder and save it to a new file

	croppedFileDir, fileName := filepath.Split(inputFilePath)
	croppedFileName := fmt.Sprintf("%scropped_%s", croppedFileDir, fileName)

	croppedFile, _ := os.Create(croppedFileName)
	defer croppedFile.Close()

	croppedFileWriter := bufio.NewWriter(croppedFile)

	// Depending on the image type, we have to choose an appropriate encoder
	switch imageType {
	case "image/png":
		err = png.Encode(croppedFileWriter, croppedImg)
	case "image/jpeg":
		err = jpeg.Encode(croppedFileWriter, croppedImg, &jpeg.Options{})
	}
	if err != nil {
		panic(err)
	}
	croppedFileWriter.Flush()
}

func main() {

	// Setup the flags
	cHeight := flag.Int("height", 0, "Crop Height")
	cWidth := flag.Int("width", 0, "Crop Width")
	flag.Parse()

	//Check flags, file name specified
	if *cHeight == 0 {
		log.Fatal("Must specify the Crop Height")
	}

	if *cWidth == 0 {
		log.Fatal("Must specify crop width")
	}

	numImages := len(flag.Args())
	if numImages == 0 {
		log.Fatal("Must specify at least 1 image to crop")
	}

	// Loop over each file , crop and save the cropped image.
	for _, inputFilePath := range flag.Args() {
		cropper(inputFilePath, *cWidth, *cHeight)
	}
}

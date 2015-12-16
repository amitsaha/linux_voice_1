/* Simple command line program to crop images

Usage:

$ go run crop.go --height=5000 --width=7000 <path to image1> <path to image 2>

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

func validateImageType(imageType string) bool {
	recognizedImageTypes := []string{"image/jpeg", "image/png"}
	for _, t := range recognizedImageTypes {
		if imageType == t {
			return true
		}
	}
	return false
}

/*
cropper() accepts three parameters:

inputFilePath : A string referring to the relative path to the image to crop

cWidth    : An int specifying the desired width of the cropped image

cHeight   : An int specifying the desired height of the cropped image
*/
func cropper(inputFilePath string, cWidth int, cHeight int) {
	imageData, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	imageType := http.DetectContentType(imageData)
	// Check if we can handle this image
	if !validateImageType(imageType) {
		log.Fatal("Cannot handle image of this type")
	}

	// We first decode the image
	reader := bytes.NewReader(imageData)
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	// Perform the cropping
	croppedImg, err := cutter.Crop(img, cutter.Config{
		Height: cHeight,
		Width:  cWidth,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Now we encode the cropped image data using the appropriate
	// encoder and save it to a new file

	croppedFileDir, fileName := filepath.Split(inputFilePath)
	croppedFileName := fmt.Sprintf("%scropped_%s", croppedFileDir, fileName)
	croppedFile, err := os.Create(croppedFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer croppedFile.Close()

	croppedFileWriter := bufio.NewWriter(croppedFile)

	// Depending on the image type, we have to choose an appropriate encoder
	if imageType == "image/png" {
		err = png.Encode(croppedFileWriter, croppedImg)
	}
	if imageType == "image/jpeg" {
		// The third argument indicates the quality of the encoding
		// We specify 100 (it can take values in [1,100], inclusive.
		err = jpeg.Encode(croppedFileWriter, croppedImg, &jpeg.Options{100})
	}
	if err != nil {
		log.Fatal(err)
	}
	croppedFileWriter.Flush()
}

func main() {

	// Setup the flags
	cHeight := flag.Int("height", 0, "Crop Height")
	cWidth := flag.Int("width", 0, "Crop Width")
	flag.Parse()

	//Check flags, file name specified
	if *cHeight <= 0 {
		log.Fatal("Must specify the crop height to be a positive integer")
	}

	if *cWidth <= 0 {
		log.Fatal("Must specify the crop width to be a positive integer")
	}

	numImages := len(flag.Args())
	if numImages == 0 {
		log.Fatal("Must specify at least 1 image to crop")
	}

	// Loop over each file, crop and save the cropped image.
	for _, inputFilePath := range flag.Args() {
		cropper(inputFilePath, *cWidth, *cHeight)
	}
}

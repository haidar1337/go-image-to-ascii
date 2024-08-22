package main

import (
	"bufio"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
)

type AsciiArtConfig struct {

	height int
	width int
}

func readImageInput() (image.Image, error) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the path to an image: ")
	scanner.Scan()

	input := scanner.Text()
	lenInput := len(input)
	format := input[lenInput - 3:lenInput]
	
	f, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if format == "png" {
		png, _ := png.Decode(f)
		return png, nil
	} else if format == "jpeg" {
		jpeg, _ := jpeg.Decode(f)
		return jpeg, nil
	}

	return nil, errors.New("an error has occurred while reading the input file")
}

func ConvertImageToGrayscale(img image.Image) image.Image {
	gray := image.NewRGBA(img.Bounds())

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			px :=  img.At(x, y)
			red, green, blue, _ := px.RGBA()
			luminosity := float64(red) * 0.299 + float64(green) * 0.587 + float64(blue) * 0.114

			gray.Set(x, y, color.Gray{
				Y: uint8(luminosity / 256),
			})
		}
	}

	return gray
}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func repl(cfg *AsciiArtConfig) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("config: %v\n", cfg.mode)
	
	for {
		fmt.Print("Enter the path to an image: ")
		scanner.Scan()

		input := scanner.Text()
		
		f, err := os.Open(input)
		defer f.Close()
		if err != nil {
			fmt.Println("Invalid image: make sure the image path is correct")
			continue
		}

		format := getImageFormat(input)
		img, err := decodeImage(f, format)
		if err != nil {
			fmt.Println(err)
			continue
		}

		gray := ConvertImageToGrayscale(img)
		asciiChars := []rune{'.', ';', '+', '*', '?', '%', 'S', '#'}
		if cfg.mode == lightMode {
			asciiChars = []rune{'#', 'S', '%', '?', '*', '+', ';', '.'}
		}
		asciiArt := MapPixels(gray, &asciiChars, cfg)

		fmt.Println(asciiArt)
	}
}

func decodeImage(file *os.File, format string) (image.Image, error) {
	var img image.Image
	var err error

	if format == "png" {
		img, err = png.Decode(file)
	} else if format == "jpeg" {
		img, err = jpeg.Decode(file)
	}

	if err != nil {
		return nil, errors.New("an error has occurred during decoding")
	}

	return img, nil
}

func getImageFormat(input string) string {
	return input[len(input) - 3:]
}
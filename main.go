package main

import (
	"fmt"
)

const (
	darkMode mode = "dark"
	lightMode mode = "light"
)

type mode string

type AsciiArtConfig struct {
	mode mode

}

func main() {
	img, err := readImageInput()
	if err != nil {
		fmt.Println(err)
	}
	cfg := AsciiArtConfig{
		mode: lightMode,
	}
	gray := ConvertImageToGrayscale(img)
	asciiChars := []rune{'.', ';', '+', '*', '?', '%', 'S', '#'}
	if cfg.mode == lightMode {
		asciiChars = []rune{'#', 'S', '%', '?', '*', '+', ';', '.'}
	}
    asciiArt := MapPixels(gray, &asciiChars, &cfg)

	fmt.Println(asciiArt)
}
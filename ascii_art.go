package main

import (
	"image"
)

func correspondingASCIICharacter(grayValue uint8, asciiChars []rune) rune {
    index := int(grayValue) * (len(asciiChars) - 1) / 255
    return asciiChars[index]
}


func MapPixels(img image.Image, asciiChars *[]rune, config *AsciiArtConfig) string {
	out := ""
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y += 8 {
		linePixels := ""
		for x := bounds.Min.X; x < bounds.Max.X; x += 4 {
			px := img.At(x, y)
			r, g, b, _ := px.RGBA()
			Y := uint8(calculateImageLuminosity(r, g, b)/256)
			
			linePixels += string(correspondingASCIICharacter(Y, *asciiChars))
		}
		out += linePixels + "\n"
	}

	return out
}
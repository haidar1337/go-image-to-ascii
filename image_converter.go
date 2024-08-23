package main

import (
	"image"
	"image/color"
)

func ConvertImageToGrayscale(img image.Image) image.Image {
	gray := image.NewRGBA(img.Bounds())

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			px :=  img.At(x, y)
			red, green, blue, _ := px.RGBA()
			luminosity := calculateImageLuminosity(red, green, blue)

			gray.Set(x, y, color.Gray{
				Y: uint8(luminosity / 256),
			})
		}
	}

	return gray
}

func calculateImageLuminosity(red, green, blue uint32) float64 {
	luminosity := float64(red) * 0.299 + float64(green) * 0.587 + float64(blue) * 0.114
	return luminosity
}
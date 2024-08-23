package main

import (
	"image"
	"image/color"
)

func resizeImage(img image.Image, height int) image.Image {
    bounds := img.Bounds()
    originalHeight := bounds.Dy()
    originalWidth := bounds.Dx()

    aspectRatio := float32(originalWidth) / float32(originalHeight)
    width := int(float32(height) * aspectRatio)
    resized := image.NewRGBA(image.Rect(0, 0, width, height))
	
    factor := float32(originalHeight) / float32(height)
    var X, Y int
    var clr color.Color
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            X = int(factor * float32(x) + 0.5)
            Y = int(factor * float32(y) + 0.5)
            clr = img.At(X, Y)
            resized.Set(x, y, clr)
        }
    }
    return resized
}
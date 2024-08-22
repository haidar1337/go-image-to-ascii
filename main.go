package main

import "fmt"


func main() {
	img, err := readImageInput()
	if err != nil {
		fmt.Println(err)
	}
	ConvertImageToGrayscale(img)
}
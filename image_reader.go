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



package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"regexp"
)

func repl(cfg *AsciiArtConfig) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("config:\nmode: %v\nscale: %v\n", cfg.mode, cfg.scale)

	for {
		fmt.Print("Enter the path to an image, or a network image URL: ")
		scanner.Scan()

		input := scanner.Text()
		matches, err := regexp.Match(`(http(s?):)([/|.|\w|\s|-])*\.(?:jpg|png)`, []byte(input))
		if err != nil {
			fmt.Println(err)
			continue
		}

		var img image.Image
		var data []byte
		if matches {
			res, err := http.Get(input)
			if err !=  nil {
				fmt.Println(err)
				continue
			}
			if res.StatusCode > 299 {
				fmt.Printf("response failed with status code %v\n", res.StatusCode)
				continue
			}
			defer res.Body.Close()

			data, err = io.ReadAll(res.Body)
			matches, err := regexp.Match(`.(?:jpg)`, []byte(input))
			if matches {
				img, err = decodeImage(data, "jpg")
				if err != nil {
					fmt.Println(err)
					continue
				}
			} else {
				img, err = decodeImage(data, "png")
				if err != nil {
					fmt.Println(err)
					continue
				}
			}

		} else {
			data, err = os.ReadFile(input)
			if err != nil {
				fmt.Println("invalid image: make sure the image path is correct")
				continue
			}

			format := getImageFormat(input)
			img, err = decodeImage(data, format)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
		
		resizedImage := resizeImage(img, int(float64(img.Bounds().Dy()) * cfg.scale))
		gray := ConvertImageToGrayscale(resizedImage)
		asciiChars := []rune{'.', ';', '+', '*', '?', '%', 'S', '#'}
		if cfg.mode == lightMode {
			asciiChars = []rune{'#', 'S', '%', '?', '*', '+', ';', '.'}
		}
		asciiArt := constructAsciiArt(gray, &asciiChars, cfg)

		fmt.Println(asciiArt)
	}
}

func decodeImage(data []byte, format string) (image.Image, error) {
	var img image.Image
	var err error

	if format == "png" {
		img, err = png.Decode(bytes.NewReader(data))
	} else if format == "jpg" {
		img, err = jpeg.Decode(bytes.NewReader(data))
	}

	if err != nil {
		return nil, err
	}

	return img, nil
}

func getImageFormat(input string) string {
	return input[len(input) - 3:]
}
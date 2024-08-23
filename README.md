# go-image-to-ascii

## Overview
This tool is used to output an ASCII art of an input image. Give it the path to your image, or paste in
a network image URL, and it will print the ASCII art of the image.

Input:
![original cat](example_input.jpg)

Output:
![ascii cat](example_output.png)

You may resize an image by configuring the scale in the config struct (0.00 to 1.00). You may also
change the mode depending on your terminal's background color (dark and light modes).

## Get Started
Go's toolchain is required to bulid the `.exe` file.
```
go build
```
An executable file will be generated, run it.
```
./go-image-to-ascii
```

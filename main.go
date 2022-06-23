package main

import (
	"bufio"
	"fmt"
	"imgres/image"
	"os"
)

func main() {
	fmt.Print("Input image URL: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	imageURL := scanner.Text()

	image := image.Image(imageURL)
	fmt.Println(image.GetResolution())
}

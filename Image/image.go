package image

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
)

type ImageResolution struct {
	ImageURL string
}

type Resolution struct {
	Width  int
	Height int
}

func (i ImageResolution) GetResolution() Resolution {
	response, err := http.Get(i.ImageURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	decodedImage, _, err := image.Decode(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	rectangle := decodedImage.Bounds()

	return Resolution{
		Width:  rectangle.Dx(),
		Height: rectangle.Dy(),
	}
}

func Image(imageURL string) ImageResolution {
	return ImageResolution{
		ImageURL: imageURL,
	}
}

package img

import (
	"image"
	"image/png"
	_ "image/png"
	"os"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func CutImage(sizeY, sizeX, posY, posX int) image.Image {
	file, err := os.Open("internal/img/spritesheet.jpg")
	if err != nil {
		panic(err)
	}
	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	rect := image.Rect(0, 0, sizeX, sizeY)
	rect = rect.Add(image.Point{posX, posY})
	return img.(SubImager).SubImage(rect)
}

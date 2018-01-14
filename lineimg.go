package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"io"

	"./line"
)

func main() {

	msg := "send an image"

	x := 0
	y := 0
	width := 100
	height := 50

	img := image.NewRGBA(image.Rect(x, y, width, height))

	b := &bytes.Buffer{}
	if err := jpeg.Encode(b, img, &jpeg.Options{100}); err != nil {
		panic(err)
	}

	if err := line.SendImage(msg, io.Reader(b)); err != nil {
		panic(err)
	}

}

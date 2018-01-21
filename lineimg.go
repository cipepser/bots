package main

import (
	"os"

	"./line"
)

func main() {
	msg := "send an image"
	filename := "./tmp.jpg"

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := line.SendImage(msg, f, filename); err != nil {
		panic(err)
	}

}

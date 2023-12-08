package gifs

import (
	"image/gif"
	"os"
)

func LoadGIF(filename string) (*gif.GIF, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return gif.DecodeAll(file)
}

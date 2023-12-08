package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"gifler/pkg/game"
	"gifler/pkg/gifs"
)

func main() {
	gifImage, err := gifs.LoadGIF("path/to/your.gif")
	if err != nil {
		log.Fatalf("Failed to load GIF: %v", err)
	}

	g := game.NewGame(gifImage)
	ebiten.SetWindowSize(gifImage.Config.Width, gifImage.Config.Height)
	ebiten.SetWindowTitle("Gifler - GIF Player")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

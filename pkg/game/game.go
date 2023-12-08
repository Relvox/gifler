package game

import (
	"image/gif"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	gifImage   *gif.GIF
	frameIndex int
}

func NewGame(gif *gif.GIF) *Game {
	return &Game{
		gifImage: gif,
	}
}

func (g *Game) Update() error {
	// Handle frame-by-frame playback logic
	// Example: Pressing 'N' key goes to the next frame
	if inpututil.IsKeyJustPressed(ebiten.KeyN) {
		g.frameIndex = (g.frameIndex + 1) % len(g.gifImage.Image)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the current frame of the GIF
	currentFrame := g.gifImage.Image[g.frameIndex]
	screen.WritePixels(currentFrame.Pix)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Return the size of the GIF as the window size
	return g.gifImage.Config.Width, g.gifImage.Config.Height
}

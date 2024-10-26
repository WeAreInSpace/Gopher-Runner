package resources

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	Player_Gopher = "./assets/gopher/gopher.png"
)

func GetImage(imagePath string) *ebiten.Image {
	image, _, newImageE := ebitenutil.NewImageFromFile(imagePath)
	if newImageE != nil {
		log.Fatal(newImageE)
	}
	return image
}

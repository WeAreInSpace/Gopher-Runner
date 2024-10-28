package resources

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	Icon16 string = "./assets/icon/16.png"
	Icon32 string = "./assets/icon/32.png"

	Gopher        string = "./assets/gopher/gopher.png"
	Player_Gopher string = "./assets/gopher/gopher_front.png"
)

func GetImage(imagePath string) *ebiten.Image {
	image, _, newImageE := ebitenutil.NewImageFromFile(imagePath)
	if newImageE != nil {
		log.Fatal(newImageE)
	}
	return image
}

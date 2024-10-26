package game

import (
	"image/color"
	"sync"

	"github.com/WeAreInSpace/Gopher-Runner/base/player"
	"github.com/WeAreInSpace/Gopher-Runner/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	Mx     *sync.Mutex
	Config *config.Config
	Player *player.Player
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Mx.Lock()
		g.Player.X -= 2
		g.Mx.Unlock()
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Mx.Lock()
		g.Player.X += 2
		g.Mx.Unlock()
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Mx.Lock()
		g.Player.Y -= 2
		g.Mx.Unlock()
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Mx.Lock()
		g.Player.Y += 2
		g.Mx.Unlock()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	imageOpts := ebiten.DrawImageOptions{}
	screen.Fill(color.RGBA{0, 150, 255, 255})

	imageOpts.GeoM.Translate(g.Player.X, g.Player.Y)
	screen.DrawImage(g.Player.Image, &imageOpts)
	imageOpts.GeoM.Reset()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth / 2, outsideHeight / 2
}

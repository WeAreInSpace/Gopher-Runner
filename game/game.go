package game

import (
	"fmt"
	"image/color"
	"net"
	"sync"

	"github.com/WeAreInSpace/Gopher-Runner/base/player"
	"github.com/WeAreInSpace/Gopher-Runner/camera"
	"github.com/WeAreInSpace/Gopher-Runner/config"
	"github.com/WeAreInSpace/Gopher-Runner/network"
	"github.com/WeAreInSpace/Gopher-Runner/packet"
	"github.com/WeAreInSpace/Gopher-Runner/resources"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	Mx *sync.Mutex
	Wg *sync.WaitGroup

	Config *config.Config
	Player *player.Player
	Camera *camera.Camera

	Conn          net.Conn
	Ib            *packet.Inbound
	Og            *packet.Outgoing
	PacketManager *network.PacketManager

	screen  *ebiten.Image
	screenW float64
	screenH float64
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Mx.Lock()
		g.Player.X -= 2
		g.PacketManager.FollowPlayer(int64(g.Player.X), int64(g.Player.Y))
		g.Mx.Unlock()
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Mx.Lock()
		g.Player.X += 2
		g.PacketManager.FollowPlayer(int64(g.Player.X), int64(g.Player.Y))
		g.Mx.Unlock()
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Mx.Lock()
		g.Player.Y -= 2
		g.PacketManager.FollowPlayer(int64(g.Player.X), int64(g.Player.Y))
		g.Mx.Unlock()
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Mx.Lock()
		g.Player.Y += 2
		g.PacketManager.FollowPlayer(int64(g.Player.X), int64(g.Player.Y))
		g.Mx.Unlock()
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) && ebiten.IsKeyPressed(ebiten.KeyControl) && !ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Mx.Lock()
		g.Player.X -= 1
		g.PacketManager.FollowPlayer(int64(g.Player.X), int64(g.Player.Y))
		g.Mx.Unlock()
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) && ebiten.IsKeyPressed(ebiten.KeyControl) && !ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Mx.Lock()
		g.Player.X += 1
		g.PacketManager.FollowPlayer(int64(g.Player.X), int64(g.Player.Y))
		g.Mx.Unlock()
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) && ebiten.IsKeyPressed(ebiten.KeyControl) && !ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Mx.Lock()
		g.Player.Y -= 1
		g.PacketManager.FollowPlayer(int64(g.Player.X), int64(g.Player.Y))
		g.Mx.Unlock()
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) && ebiten.IsKeyPressed(ebiten.KeyShift) {
		g.Mx.Lock()
		g.Player.X += 1.5
		g.PacketManager.FollowPlayer(int64(g.Player.X), int64(g.Player.Y))
		g.Mx.Unlock()
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) && ebiten.IsKeyPressed(ebiten.KeyShift) {
		g.Mx.Lock()
		g.Player.X -= 1.5
		g.PacketManager.FollowPlayer(int64(g.Player.X), int64(g.Player.Y))
		g.Mx.Unlock()
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) && ebiten.IsKeyPressed(ebiten.KeyShift) {
		g.Mx.Lock()
		g.Player.Y += 1.5
		g.PacketManager.FollowPlayer(int64(g.Player.X), int64(g.Player.Y))
		g.Mx.Unlock()
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) && ebiten.IsKeyPressed(ebiten.KeyShift) {
		g.Mx.Lock()
		g.Player.Y -= 1.5
		g.PacketManager.FollowPlayer(int64(g.Player.X), int64(g.Player.Y))
		g.Mx.Unlock()
	}

	g.Camera.FollowTarget(g.Player.X+16, g.Player.Y+16, g.screenW, g.screenH)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.screen = screen
	screen.Fill(color.RGBA{65, 201, 226, 255})
	imageOpts := ebiten.DrawImageOptions{}

	imageOpts.GeoM.Translate(0, 0)
	imageOpts.GeoM.Translate(g.Camera.X, g.Camera.Y)
	screen.DrawImage(resources.GetImage(resources.Gopher), &imageOpts)
	imageOpts.GeoM.Reset()

	imageOpts.GeoM.Translate(g.Player.X, g.Player.Y)
	imageOpts.GeoM.Translate(g.Camera.X, g.Camera.Y)
	screen.DrawImage(g.Player.Image, &imageOpts)
	imageOpts.GeoM.Reset()

	ebitenutil.DebugPrint(screen, fmt.Sprintf("%4.f", ebiten.ActualFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	g.screenW = float64(outsideWidth / 4)
	g.screenH = float64(outsideHeight / 4)
	return outsideWidth / 4, outsideHeight / 4
}

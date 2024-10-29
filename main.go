package main

import (
	"image"
	"log"
	"sync"

	"github.com/WeAreInSpace/Gopher-Runner/base/player"
	"github.com/WeAreInSpace/Gopher-Runner/base/sprite"
	"github.com/WeAreInSpace/Gopher-Runner/camera"
	"github.com/WeAreInSpace/Gopher-Runner/config"
	"github.com/WeAreInSpace/Gopher-Runner/game"
	"github.com/WeAreInSpace/Gopher-Runner/network"
	"github.com/WeAreInSpace/Gopher-Runner/resources"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	log.Println("Loading library, please wait.")
	var mx sync.Mutex
	var wg sync.WaitGroup

	log.Println("Loading game config.")
	gormConfig := gorm.Config{
		QueryFields: true,
	}
	configDb, openConfigE := gorm.Open(
		sqlite.Open("config.db"),
		&gormConfig,
	)
	if openConfigE != nil {
		log.Fatal(openConfigE)
	}
	configDb.AutoMigrate(&config.ConfigSchema{})

	err := configDb.First(&config.ConfigSchema{}).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		log.Println("Createing default config.")
		defaultConfig := &config.ConfigSchema{
			Fullscreen:  false,
			AlwaysOnTop: false,

			PlayerSound:  100,
			MusicSound:   100,
			AmbientSound: 100,
		}
		configDb.Create(defaultConfig)
	}

	config := &config.Config{
		Mx: &mx,
		Db: configDb,
	}
	log.Println("Loaded game config.")

	log.Println("Setting game window.")
	ebiten.SetWindowSize(1200, 720)
	ebiten.SetWindowTitle("-=GOPHER RUNNER")
	ebiten.SetFullscreen(config.GetFullScreen())
	ebiten.SetWindowResizingMode(
		ebiten.WindowResizingModeEnabled,
	)

	_, iconImage16, newIconImage16E := ebitenutil.NewImageFromFile(resources.Icon16)
	if newIconImage16E != nil {
		log.Fatal(newIconImage16E)
	}
	_, iconImage32, newIconImage32E := ebitenutil.NewImageFromFile(resources.Icon32)
	if newIconImage32E != nil {
		log.Fatal(newIconImage32E)
	}
	icon := []image.Image{}
	icon = append(icon, iconImage16)
	icon = append(icon, iconImage32)
	ebiten.SetWindowIcon(icon)
	log.Println("Set game window successfully.")

	log.Println("Loading game struct")

	conn, ib, og := network.HandleConn(":25201")
	packetManager := network.PacketManager{
		Conn: conn,
		Ib:   &ib,
		Og:   &og,
	}

	playerHandshakeEvent := make(chan string)

	log.Println("Reloading MOTD")
	packetManager.GetMOTD()

	go func() {
		handshakeE := packetManager.Handshake(
			network.PlayerHandshake{
				Name: "TEST Player",
				Uuid: "weareinspace",
			},
			playerHandshakeEvent,
		)
		if handshakeE != nil {
			log.Fatal(handshakeE)
		}
	}()

	log.Printf("EVENT: Player hanshake %s", <-playerHandshakeEvent)

	playerGopher := resources.GetImage(resources.Player_Gopher)
	game := game.Game{
		Mx: &mx,
		Wg: &wg,

		Config: config,

		PacketManager: &packetManager,
		Conn:          conn,
		Ib:            &ib,
		Og:            &og,

		Camera: camera.NewCamera(0.0, 0.0),

		Player: &player.Player{
			Sprite: &sprite.Sprite{
				Image: playerGopher,
				X:     0,
				Y:     0,
			},
			Health: 10,
		},
	}
	log.Println("Loaded game struct")
	log.Println("The game will be in your screen. :)")

	runGameE := ebiten.RunGame(&game)
	if runGameE != nil {
		log.Fatal(runGameE)
	}
}

package main

import (
	"log"
	"sync"

	"github.com/WeAreInSpace/Gopher-Runner/base/player"
	"github.com/WeAreInSpace/Gopher-Runner/base/sprite"
	"github.com/WeAreInSpace/Gopher-Runner/config"
	"github.com/WeAreInSpace/Gopher-Runner/game"
	"github.com/WeAreInSpace/Gopher-Runner/resources"

	"github.com/hajimehoshi/ebiten/v2"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var mx sync.Mutex

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
		log.Default().Println("Createing default config")
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

	ebiten.SetWindowSize(1200, 720)
	ebiten.SetWindowTitle("-=GOPHER RUNNER")
	ebiten.SetFullscreen(config.GetFullScreen())
	ebiten.SetWindowResizingMode(
		ebiten.WindowResizingModeEnabled,
	)

	playerGopher := resources.GetImage(resources.Player_Gopher)

	game := game.Game{
		Mx:     &mx,
		Config: config,
		Player: &player.Player{
			Sprite: &sprite.Sprite{
				Image: playerGopher,
				X:     0,
				Y:     0,
			},
			Health: 2,
		},
	}
	runGameE := ebiten.RunGame(&game)

	if runGameE != nil {
		log.Fatal(runGameE)
	}
}

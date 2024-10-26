package config

import (
	"log"
	"sync"

	"gorm.io/gorm"
)

type ConfigSchema struct {
	gorm.Model

	Fullscreen  bool
	AlwaysOnTop bool

	PlayerSound  uint8
	MusicSound   uint8
	AmbientSound uint8
}

type Config struct {
	Mx *sync.Mutex

	Db *gorm.DB
}

/*
```` Fullscreen
*/

func (c *Config) GetFullScreen() bool {
	c.Mx.Lock()
	var config ConfigSchema
	result := c.Db.First(&config)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	defer c.Mx.Unlock()
	return config.Fullscreen
}

func (c *Config) SetFullScreen(newValue bool) {
	c.Mx.Lock()
	var config ConfigSchema
	result := c.Db.Model(&config).Update("fullscreen", newValue)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	defer c.Mx.Unlock()
}

/*
```` alwaysOnTop
*/

func (c *Config) GetAlwaysOnTop() bool {
	c.Mx.Lock()
	var config ConfigSchema
	result := c.Db.First(&config, "always_on_top")
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	defer c.Mx.Unlock()
	return config.AlwaysOnTop
}

func (c *Config) SetAlwaysOnTop(newValue bool) {
	c.Mx.Lock()
	var config ConfigSchema
	result := c.Db.Model(&config).Update("always_on_top", newValue)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	defer c.Mx.Unlock()
}

/*
```` playerSound
*/

func (c *Config) GetPlayerSound() uint8 {
	c.Mx.Lock()
	var config ConfigSchema
	result := c.Db.First(&config, "player_sound")
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	defer c.Mx.Unlock()
	return config.PlayerSound
}

func (c *Config) SetPlayerSound(newValue uint8) {
	c.Mx.Lock()
	var config ConfigSchema
	result := c.Db.Model(&config).Update("player_sound", newValue)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	defer c.Mx.Unlock()
}

/*
```` musicSound
*/

func (c *Config) GetMusicSound() uint8 {
	c.Mx.Lock()
	var config ConfigSchema
	result := c.Db.First(&config, "music_sound")
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	defer c.Mx.Unlock()
	return config.MusicSound
}

func (c *Config) SetMusicSound(newValue uint8) {
	c.Mx.Lock()
	var config ConfigSchema
	result := c.Db.Model(&config).Update("music_sound", newValue)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	defer c.Mx.Unlock()
}

/*
```` ambientSound
*/

func (c *Config) GetAmbientSound() uint8 {
	c.Mx.Lock()
	var config ConfigSchema
	result := c.Db.First(&config, "ambient_sound")
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	defer c.Mx.Unlock()
	return config.AmbientSound
}

func (c *Config) SetAmbientSound(newValue uint8) {
	c.Mx.Lock()
	var config ConfigSchema
	result := c.Db.Model(&config).Update("ambient_sound", newValue)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	defer c.Mx.Unlock()
}

package entity

import "github.com/WeAreInSpace/Gopher-Runner/base/sprite"

type Entity struct {
	*sprite.Sprite

	isHostileMobs bool
	ai            bool
}

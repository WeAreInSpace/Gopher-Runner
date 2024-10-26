package player

import "github.com/WeAreInSpace/Gopher-Runner/base/sprite"

type Player struct {
	*sprite.Sprite

	Health int8
}

package camera

type Camera struct {
	X, Y float64
}

func NewCamera(x, y float64) *Camera {
	return &Camera{
		X: x,
		Y: y,
	}
}

func (c *Camera) FollowTarget(tgX, tgY, screenW, screenH float64) {
	c.X = -tgX + screenW/2.0
	c.Y = -tgY + screenH/2.0
}

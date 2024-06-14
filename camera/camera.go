package camera

import (
	"three-go/math/vector"
)

type Camera struct {
	Position vector.Vector3
	Rotation vector.Vector3
}

func NewCamera() *Camera {
	return &Camera{
		Position: vector.Vector3{X: 0, Y: 0, Z: 0},
		Rotation: vector.Vector3{X: 0, Y: 0, Z: 0},
	}
}

func (c *Camera) MoveForward(amount float64) {
	c.Position.Z -= amount
}

func (c *Camera) MoveBackward(amount float64) {
	c.Position.Z += amount
}

func (c *Camera) MoveLeft(amount float64) {
	c.Position.X -= amount
}

func (c *Camera) MoveRight(amount float64) {
	c.Position.X += amount
}

func (c *Camera) MoveUp(amount float64) {
	c.Position.Y += amount
}

func (c *Camera) MoveDown(amount float64) {
	c.Position.Y -= amount
}

func (c *Camera) RotateY(angle float64) {
	c.Rotation.Y += angle
}

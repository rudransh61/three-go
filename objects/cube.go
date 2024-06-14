package objects

import (
	"three-go/math/vector"
)

type Cube struct {
	Vertices []vector.Vector3
	Position vector.Vector3
	Rotation vector.Vector3
}

func NewCube(size float64) *Cube {
	halfSize := size / 2
	vertices := []vector.Vector3{
		{X: -halfSize, Y: -halfSize, Z: -halfSize},
		{X: halfSize, Y: -halfSize, Z: -halfSize},
		{X: halfSize, Y: halfSize, Z: -halfSize},
		{X: -halfSize, Y: halfSize, Z: -halfSize},
		{X: -halfSize, Y: -halfSize, Z: halfSize},
		{X: halfSize, Y: -halfSize, Z: halfSize},
		{X: halfSize, Y: halfSize, Z: halfSize},
		{X: -halfSize, Y: halfSize, Z: halfSize},
	}
	return &Cube{
		Vertices: vertices,
		Position: vector.Vector3{X: 0, Y: 0, Z: 0},
		Rotation: vector.Vector3{X: 0, Y: 0, Z: 0},
	}
}

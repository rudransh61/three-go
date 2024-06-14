package camera

import "three-go/math/vector"

type Camera struct {
	Position vector.Vector3
	Target   vector.Vector3
}

func NewCamera() *Camera {
	return &Camera{
		Position: vector.Vector3{X: 0, Y: 0, Z: 0},
		Target:   vector.Vector3{X: 0, Y: 0, Z: -1},
	}
}

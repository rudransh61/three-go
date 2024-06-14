package objects

import "three-go/math/vector"

type Object3D struct {
	Position vector.Vector3
	Rotation vector.Vector3
	Scale    vector.Vector3
	Vertices []vector.Vector3
}

func NewObject3D() *Object3D {
	return &Object3D{
		Position: vector.Vector3{X: 0, Y: 0, Z: 0},
		Rotation: vector.Vector3{X: 0, Y: 0, Z: 0},
		Scale:    vector.Vector3{X: 1, Y: 1, Z: 1},
		Vertices: []vector.Vector3{},
	}
}

func NewCube(size float64) *Object3D {
	half := size / 2
	vertices := []vector.Vector3{
		{-half, -half, -half},
		{half, -half, -half},
		{half, half, -half},
		{-half, half, -half},
		{-half, -half, half},
		{half, -half, half},
		{half, half, half},
		{-half, half, half},
	}
	return &Object3D{
		Position: vector.Vector3{X: 0, Y: 0, Z: 0},
		Rotation: vector.Vector3{X: 0, Y: 0, Z: 0},
		Scale:    vector.Vector3{X: 1, Y: 1, Z: 1},
		Vertices: vertices,
	}
}

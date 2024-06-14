package objects

import (
	"image/color"
	"three-go/math/vector"
)

type Triangle struct {
	Vertices [3]vector.Vector3
	Color    color.RGBA
}

type TriangleMesh struct {
	Position  vector.Vector3
	Rotation  vector.Vector3
	Scale     vector.Vector3
	Triangles []Triangle
}

func NewTriangleMesh() *TriangleMesh {
	return &TriangleMesh{
		Position:  vector.Vector3{X: 0, Y: 0, Z: 0},
		Rotation:  vector.Vector3{X: 0, Y: 0, Z: 0},
		Scale:     vector.Vector3{X: 1, Y: 1, Z: 1},
		Triangles: []Triangle{},
	}
}

func NewColoredTriangleMesh() *TriangleMesh {
	triangles := []Triangle{
		{
			Vertices: [3]vector.Vector3{
				{X: -1, Y: -1, Z: 0},
				{X: 1, Y: -1, Z: 0},
				{X: 0, Y: 1, Z: 0},
			},
			Color: color.RGBA{R: 255, G: 0, B: 0, A: 255},
		},
		{
			Vertices: [3]vector.Vector3{
				{X: -1, Y: -1, Z: -1},
				{X: 1, Y: -1, Z: -1},
				{X: 0, Y: 1, Z: -1},
			},
			Color: color.RGBA{R: 0, G: 255, B: 0, A: 255},
		},
	}
	return &TriangleMesh{
		Position:  vector.Vector3{X: 0, Y: 0, Z: 0},
		Rotation:  vector.Vector3{X: 0, Y: 0, Z: 0},
		Scale:     vector.Vector3{X: 1, Y: 1, Z: 1},
		Triangles: triangles,
	}
}

package main

import (
	"image/color"
	"log"
	"math"

	"three-go/camera"
	"three-go/math/vector"
	"three-go/objects"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Game struct {
	camera *camera.Camera
	cube   *objects.Object3D
}

func (g *Game) Update() error {
	// Update game logic here
	g.cube.Rotation.Y += 0.01
	if g.cube.Rotation.Y > 2*math.Pi {
		g.cube.Rotation.Y -= 2 * math.Pi
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear screen with a background color
	screen.Fill(color.RGBA{0, 0, 0, 255})

	// Draw the cube edges
	edges := [][2]int{
		{0, 1}, {1, 2}, {2, 3}, {3, 0},
		{4, 5}, {5, 6}, {6, 7}, {7, 4},
		{0, 4}, {1, 5}, {2, 6}, {3, 7},
	}

	for _, edge := range edges {
		start := g.cube.Vertices[edge[0]]
		end := g.cube.Vertices[edge[1]]

		// Apply rotation
		rotatedStart := rotateVertex(start, g.cube.Rotation)
		rotatedEnd := rotateVertex(end, g.cube.Rotation)

		// Apply translation
		translatedStart := vector.Vector3{
			X: rotatedStart.X + g.cube.Position.X,
			Y: rotatedStart.Y + g.cube.Position.Y,
			Z: rotatedStart.Z + g.cube.Position.Z,
		}
		translatedEnd := vector.Vector3{
			X: rotatedEnd.X + g.cube.Position.X,
			Y: rotatedEnd.Y + g.cube.Position.Y,
			Z: rotatedEnd.Z + g.cube.Position.Z,
		}

		// Project 3D points to 2D space
		x1, y1 := project(translatedStart, g.camera)
		x2, y2 := project(translatedEnd, g.camera)

		// Draw the edge
		ebitenutil.DrawLine(screen, x1, y1, x2, y2, color.White)
	}

	// Draw some text for demonstration
	ebitenutil.DebugPrint(screen, "Three-Go with Ebiten")
}

func rotateVertex(v, rotation vector.Vector3) vector.Vector3 {
	// Rotate around Y axis for simplicity
	sinY, cosY := math.Sincos(rotation.Y)

	x := v.X*cosY - v.Z*sinY
	z := v.X*sinY + v.Z*cosY

	return vector.Vector3{X: x, Y: v.Y, Z: z}
}

func project(v vector.Vector3, cam *camera.Camera) (float64, float64) {
	// Simple perspective projection
	fov := 1.0 / math.Tan(90.0/2.0)
	aspect := float64(screenWidth) / float64(screenHeight)
	z := v.Z - cam.Position.Z
	x := (v.X - cam.Position.X) * fov / z * aspect
	y := (v.Y - cam.Position.Y) * fov / z

	// Convert to screen coordinates
	screenX := (x + 1) * screenWidth / 2
	screenY := (-y + 1) * screenHeight / 2

	return screenX, screenY
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	// Initialize camera and cube
	cam := camera.NewCamera()
	cam.Position.Z = 5
	cube := objects.NewCube(2)

	game := &Game{
		camera: cam,
		cube:   cube,
	}

	// Run the Ebiten game loop
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Three-Go with Ebiten")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

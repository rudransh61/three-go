package main

import (
	"image/color"
	"log"
	"math"

	"three-go/camera"
	"three-go/math/matrix"
	"three-go/math/vector"
	"three-go/objects"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 800
	screenHeight = 600
	rotateSpeed  = 0.01
)

type Game struct {
	camera *camera.Camera
	cube   *objects.Cube
}

func (g *Game) Update() error {
	// Rotate cube around main axes
	g.cube.Rotation.X += rotateSpeed
	g.cube.Rotation.Y += rotateSpeed
	g.cube.Rotation.Z += rotateSpeed

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear screen with a background color
	screen.Fill(color.RGBA{0, 0, 0, 255})

	// Draw lines between cube edges
	drawCubeEdges(screen, g.cube, g.camera)

	// Draw X, Y, Z axes lines
	drawAxes(screen, g.camera)

	// Draw some text for demonstration
	ebitenutil.DebugPrint(screen, "Cube Rotation Example")
}

func drawCubeEdges(screen *ebiten.Image, cube *objects.Cube, cam *camera.Camera) {
	// Project cube vertices and draw lines between edges
	// fov := 1.0 / math.Tan(math.Pi/4) // 45 degree field of view
	aspect := float64(screenWidth) / float64(screenHeight)

	for i := 0; i < len(cube.Vertices); i++ {
		v0 := cube.Vertices[i]

		// Apply rotation
		rotX := matrix.RotationX(cube.Rotation.X)
		rotY := matrix.RotationY(cube.Rotation.Y)
		rotZ := matrix.RotationZ(cube.Rotation.Z)

		rotated := matrix.MultiplyVector(rotX, v0)
		rotated = matrix.MultiplyVector(rotY, rotated)
		rotated = matrix.MultiplyVector(rotZ, rotated)

		// Apply translation
		translated := vector.Vector3{
			X: rotated.X + cube.Position.X,
			Y: rotated.Y + cube.Position.Y,
			Z: rotated.Z + cube.Position.Z,
		}

		// Project 3D point to 2D space
		x0, y0 := project(translated, cam)

		for j := i + 1; j < len(cube.Vertices); j++ {
			v1 := cube.Vertices[j]

			// Apply rotation
			rotated = matrix.MultiplyVector(rotX, v1)
			rotated = matrix.MultiplyVector(rotY, rotated)
			rotated = matrix.MultiplyVector(rotZ, rotated)

			// Apply translation
			translated = vector.Vector3{
				X: rotated.X + cube.Position.X,
				Y: rotated.Y + cube.Position.Y,
				Z: rotated.Z + cube.Position.Z,
			}

			// Project 3D point to 2D space
			x1, y1 := project(translated, cam)

			// Draw line between vertices
			ebitenutil.DrawLine(screen, x0, y0, x1, y1, color.White)

			// Update starting point of next line
			x0, y0 = x1, y1
		}
	}
}

func drawAxes(screen *ebiten.Image, cam *camera.Camera) {
	// Draw X, Y, Z axes lines
	axesLength := 3.0 // Length of axes lines

	// X-axis (red)
	ebitenutil.DrawLine(screen, 0, screenHeight/2, screenWidth, screenHeight/2, color.RGBA{255, 0, 0, 255})

	// Y-axis (green)
	ebitenutil.DrawLine(screen, screenWidth/2, 0, screenWidth/2, screenHeight, color.RGBA{0, 255, 0, 255})

	// Z-axis (blue)
	// Project 3D points to 2D space
	x, y := project(vector.Vector3{X: 0, Y: 0, Z: -axesLength}, cam)
	x0, y0 := project(vector.Vector3{X: 0, Y: 0, Z: axesLength}, cam)
	ebitenutil.DrawLine(screen, x, y, x0, y0, color.RGBA{0, 0, 255, 255})
}

func project(v vector.Vector3, cam *camera.Camera) (float64, float64) {
	// Simple perspective projection
	fov := 1.0 / math.Tan(math.Pi/4) // 45 degree field of view
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
	ebiten.SetWindowTitle("Cube Rotation Example")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

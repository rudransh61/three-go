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
	moveSpeed    = 0.1
	rotateSpeed  = 0.05
)

type Game struct {
	camera       *camera.Camera
	triangleMesh *objects.TriangleMesh
}

func (g *Game) Update() error {
	// Handle camera movement
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.camera.MoveForward(moveSpeed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.camera.MoveBackward(moveSpeed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.camera.MoveLeft(moveSpeed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.camera.MoveRight(moveSpeed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		g.camera.MoveUp(moveSpeed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		g.camera.MoveDown(moveSpeed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.camera.RotateY(-rotateSpeed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.camera.RotateY(rotateSpeed)
	}

	// Update game logic here
	g.triangleMesh.Rotation.Y += 0.01
	if g.triangleMesh.Rotation.Y > 2*math.Pi {
		g.triangleMesh.Rotation.Y -= 2 * math.Pi
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear screen with a background color
	screen.Fill(color.RGBA{0, 0, 0, 255})

	// Draw the triangle mesh
	for _, triangle := range g.triangleMesh.Triangles {
		var projectedVertices [3]vector.Vector3
		for i, vertex := range triangle.Vertices {
			// Apply rotation
			rotated := rotateVertex(vertex, g.triangleMesh.Rotation)

			// Apply translation
			translated := vector.Vector3{
				X: rotated.X + g.triangleMesh.Position.X,
				Y: rotated.Y + g.triangleMesh.Position.Y,
				Z: rotated.Z + g.triangleMesh.Position.Z,
			}

			// Project 3D point to 2D space
			x, y := project(translated, g.camera)

			// Store the projected vertices
			projectedVertices[i] = vector.Vector3{X: x, Y: y, Z: 0}
		}

		// Draw the triangle
		ebitenutil.DrawLine(screen, projectedVertices[0].X, projectedVertices[0].Y, projectedVertices[1].X, projectedVertices[1].Y, triangle.Color)
		ebitenutil.DrawLine(screen, projectedVertices[1].X, projectedVertices[1].Y, projectedVertices[2].X, projectedVertices[2].Y, triangle.Color)
		ebitenutil.DrawLine(screen, projectedVertices[2].X, projectedVertices[2].Y, projectedVertices[0].X, projectedVertices[0].Y, triangle.Color)
	}

	// Draw some text for demonstration
	ebitenutil.DebugPrint(screen, "Three-Go with Ebiten")
}

func rotateVertex(v, rotation vector.Vector3) vector.Vector3 {
	rotX := matrix.RotationX(rotation.X)
	rotY := matrix.RotationY(rotation.Y)
	rotZ := matrix.RotationZ(rotation.Z)

	v = matrix.MultiplyVector(rotX, v)
	v = matrix.MultiplyVector(rotY, v)
	v = matrix.MultiplyVector(rotZ, v)

	return v
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
	// Initialize camera and triangle mesh
	cam := camera.NewCamera()
	cam.Position.Z = 5
	triangleMesh := objects.NewColoredTriangleMesh()

	game := &Game{
		camera:       cam,
		triangleMesh: triangleMesh,
	}

	// Run the Ebiten game loop
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Three-Go with Ebiten")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

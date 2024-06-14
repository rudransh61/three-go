package main

import (
	"image/color"
	"log"

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
	moveSpeed    = 0.1
)

type Game struct {
	camera *camera.Camera
	cube   *objects.Cube
}

func (g *Game) Update() error {
	// Handle camera controls
	handleCameraControls(g.camera, g.cube)

	// Rotate cube around main axes
	// g.cube.Rotation.X += rotateSpeed
	g.cube.Rotation.Y += rotateSpeed
	// g.cube.Rotation.Z += rotateSpeed

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear screen with a background color
	screen.Fill(color.RGBA{0, 0, 0, 255})

	// Draw the cube edges
	drawCubeEdges(screen, g.cube, g.camera)

	// Draw axes
	drawAxes(screen, g.camera)

	// Draw some text for demonstration
	ebitenutil.DebugPrint(screen, "Cube Rotation Example\nArrow Keys to Move Camera\nQ/E, W/S, A/D to Rotate Camera")
}

func drawCubeEdges(screen *ebiten.Image, cube *objects.Cube, cam *camera.Camera) {
	// Edges defined by pairs of indices
	edges := [][]int{
		{0, 1}, {1, 2}, {2, 3}, {3, 0}, // Bottom face
		{4, 5}, {5, 6}, {6, 7}, {7, 4}, // Top face
		{0, 4}, {1, 5}, {2, 6}, {3, 7}, // Connecting edges
	}

	for _, edge := range edges {
		v1 := cube.Vertices[edge[0]]
		v2 := cube.Vertices[edge[1]]

		// Apply rotation
		rotX := matrix.RotationX(cube.Rotation.X)
		rotY := matrix.RotationY(cube.Rotation.Y)
		rotZ := matrix.RotationZ(cube.Rotation.Z)

		v1 = matrix.MultiplyVector(rotX, v1)
		v1 = matrix.MultiplyVector(rotY, v1)
		v1 = matrix.MultiplyVector(rotZ, v1)

		v2 = matrix.MultiplyVector(rotX, v2)
		v2 = matrix.MultiplyVector(rotY, v2)
		v2 = matrix.MultiplyVector(rotZ, v2)

		// Apply translation
		v1 = vector.Vector3{
			X: v1.X + cube.Position.X,
			Y: v1.Y + cube.Position.Y,
			Z: v1.Z + cube.Position.Z,
		}

		v2 = vector.Vector3{
			X: v2.X + cube.Position.X,
			Y: v2.Y + cube.Position.Y,
			Z: v2.Z + cube.Position.Z,
		}

		// Project 3D points to 2D space (orthographic projection)
		x1, y1 := orthographicProjection(v1, cam)
		x2, y2 := orthographicProjection(v2, cam)

		// Draw the edge
		ebitenutil.DrawLine(screen, x1, y1, x2, y2, color.White)
	}
}

func drawAxes(screen *ebiten.Image, cam *camera.Camera) {
	// Axis lengths for X, Y, Z
	axisLength := 100.0

	// X-axis (red)
	xAxisEnd := vector.Vector3{X: -axisLength, Y: 0, Z: 0}
	xAxisEnd = matrix.RotateVector(xAxisEnd, cam.Rotation)
	xAxisEnd.X += cam.Position.X
	xAxisEnd.Y += cam.Position.Y
	xAxisEnd.Z += cam.Position.Z
	x, y := orthographicProjection(xAxisEnd, cam)
	ebitenutil.DrawLine(screen, screenWidth/2, screenHeight/2, x, y, color.RGBA{255, 0, 0, 255})

	// Y-axis (green)
	yAxisEnd := vector.Vector3{X: 0, Y: -axisLength, Z: 0}
	yAxisEnd = matrix.RotateVector(yAxisEnd, cam.Rotation)
	yAxisEnd.X += cam.Position.X
	yAxisEnd.Y += cam.Position.Y
	yAxisEnd.Z += cam.Position.Z
	x, y = orthographicProjection(yAxisEnd, cam)
	ebitenutil.DrawLine(screen, screenWidth/2, screenHeight/2, x, y, color.RGBA{0, 255, 0, 255})

	// Z-axis (blue)
	zAxisEnd := vector.Vector3{X: 50, Y: 50, Z: -axisLength}
	zAxisEnd = matrix.RotateVector(zAxisEnd, cam.Rotation)
	zAxisEnd.X += cam.Position.X
	zAxisEnd.Y += cam.Position.Y
	zAxisEnd.Z += cam.Position.Z
	x, y = orthographicProjection(zAxisEnd, cam)
	ebitenutil.DrawLine(screen, screenWidth/2, screenHeight/2, x, y, color.RGBA{0, 0, 255, 255})
}

func orthographicProjection(v vector.Vector3, cam *camera.Camera) (float64, float64) {
	// Orthographic projection
	x := v.X - cam.Position.X + float64(screenWidth)/2
	y := v.Y - cam.Position.Y + float64(screenHeight)/2

	return x, y
}

func handleCameraControls(cam *camera.Camera, cube *objects.Cube) {
	// Handle camera movement and rotation controls
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		cam.Position.Y -= moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		cam.Position.Y += moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		cam.Position.X -= moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		cam.Position.X += moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		cube.Rotation.X += rotateSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		cube.Rotation.X -= rotateSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		cube.Rotation.Y += rotateSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		cube.Rotation.Y -= rotateSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		cube.Rotation.Z += rotateSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		cube.Rotation.Z -= rotateSpeed
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	// Initialize camera and cube
	cam := camera.NewCamera()
	cam.Position.Z = 0
	cube := objects.NewCube(50)
	cube.Position.X += 100

	game := &Game{
		camera: cam,
		cube:   cube,
	}

	// Run the Ebiten game loop
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Cube Rotation Example (Orthographic)")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

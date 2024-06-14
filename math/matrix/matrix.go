package matrix

import (
	"math"
	"three-go/math/vector"
)

type Matrix4 [4][4]float64

func Identity() Matrix4 {
	return Matrix4{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

func RotationX(angle float64) Matrix4 {
	c, s := math.Cos(angle), math.Sin(angle)
	return Matrix4{
		{1, 0, 0, 0},
		{0, c, -s, 0},
		{0, s, c, 0},
		{0, 0, 0, 1},
	}
}

func RotationY(angle float64) Matrix4 {
	c, s := math.Cos(angle), math.Sin(angle)
	return Matrix4{
		{c, 0, s, 0},
		{0, 1, 0, 0},
		{-s, 0, c, 0},
		{0, 0, 0, 1},
	}
}

func RotationZ(angle float64) Matrix4 {
	c, s := math.Cos(angle), math.Sin(angle)
	return Matrix4{
		{c, -s, 0, 0},
		{s, c, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

func MultiplyMatrix(a, b Matrix4) Matrix4 {
	var result Matrix4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i][j] = 0
			for k := 0; k < 4; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return result
}

func MultiplyVector(m Matrix4, v vector.Vector3) vector.Vector3 {
	return vector.Vector3{
		X: m[0][0]*v.X + m[0][1]*v.Y + m[0][2]*v.Z + m[0][3],
		Y: m[1][0]*v.X + m[1][1]*v.Y + m[1][2]*v.Z + m[1][3],
		Z: m[2][0]*v.X + m[2][1]*v.Y + m[2][2]*v.Z + m[2][3],
	}
}

// RotateVector rotates a 3D vector around the camera's rotation angles (X, Y, Z).
func RotateVector(v vector.Vector3, rotation vector.Vector3) vector.Vector3 {
	// Create rotation matrices for X, Y, Z rotations
	rotX := RotationX(rotation.X)
	rotY := RotationY(rotation.Y)
	rotZ := RotationZ(rotation.Z)

	// Apply rotations in sequence: Z, Y, X
	result := MultiplyVector(rotZ, v)
	result = MultiplyVector(rotY, result)
	result = MultiplyVector(rotX, result)

	return result
}

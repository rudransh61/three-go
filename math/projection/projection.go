package projection

import "math"

func Perspective(fov, aspect, near, far float64) Matrix4 {
	f := 1.0 / math.Tan(fov/2)
	nf := 1 / (near - far)

	return Matrix4{
		f / aspect, 0, 0, 0,
		0, f, 0, 0,
		0, 0, (far + near) * nf, -1,
		0, 0, (2 * far * near) * nf, 0,
	}
}

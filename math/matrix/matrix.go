package matrix

type Matrix4 [16]float64

func NewMatrix4() Matrix4 {
	return Matrix4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func (m *Matrix4) Multiply(other Matrix4) Matrix4 {
	var result Matrix4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i*4+j] = 0
			for k := 0; k < 4; k++ {
				result[i*4+j] += m[i*4+k] * other[k*4+j]
			}
		}
	}
	return result
}

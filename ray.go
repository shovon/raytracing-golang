package main

// Ray represents a ray
type Ray struct {
	A Vec3
	B Vec3
}

// Origin gets the origin (vector A)
func (r Ray) Origin() Vec3 {
	return r.A
}

// Direction gets the direction (vector B)
func (r Ray) Direction() Vec3 {
	return r.B
}

// PointAtParameter gets the point at parameter
func (r Ray) PointAtParameter(t float32) Vec3 {
	return r.A.Add(r.B.ScalarMultiply(t))
}

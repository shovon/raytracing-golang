package main

// Camera represents a camera.
type Camera struct {
	Origin          Vec3
	Horizontal      Vec3
	Vertical        Vec3
	LowerLeftCorner Vec3
}

// NewCamera gets a camera with some defaults
func NewCamera() Camera {
	return Camera{
		Vec3{0.0, 0.0, 0.0},
		Vec3{4.0, 0.0, 0.0},
		Vec3{0.0, 2.0, 0.0},
		Vec3{-2.0, -1.0, -1.0},
	}
}

// GetRay gets the ray from the camera.
func (c Camera) GetRay(u, v float32) Ray {
	return Ray{
		c.Origin,
		c.LowerLeftCorner.
			Add(c.Horizontal.ScalarMultiply(u)).
			Add(c.Vertical.ScalarMultiply(v)).
			Subtract(c.Origin),
	}
}

package main

import "github.com/chewxy/math32"

// Camera represents a camera.
type Camera struct {
	Origin          Vec3
	Horizontal      Vec3
	Vertical        Vec3
	LowerLeftCorner Vec3
}

// NewCamera gets a camera with some defaults
func NewCamera(lookFrom, lookAt, up Vec3, fov, aspect float32) Camera {
	var u, v, w Vec3

	var camera Camera

	theta := fov * math32.Pi / 180
	halfHeight := math32.Tan(theta / 2)
	halfWidth := aspect * halfHeight

	camera.Origin = lookFrom

	w = lookFrom.Subtract(lookAt).Normalized()
	u = up.Cross(w).Normalized()
	v = w.Cross(u)

	// camera.LowerLeftCorner = Vec3{-halfWidth, -halfHeight, -1.0}
	camera.LowerLeftCorner = camera.Origin.
		Subtract(u.ScalarMultiply(halfWidth)).
		Subtract(v.ScalarMultiply(halfHeight)).
		Subtract(w)
	camera.Horizontal = u.ScalarMultiply(2.0 * halfWidth)
	camera.Vertical = v.ScalarMultiply(2.0 * halfHeight)

	return camera
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

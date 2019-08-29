package main

import "github.com/chewxy/math32"

// Camera represents a camera.
type Camera struct {
	Origin          Vec3
	Horizontal      Vec3
	Vertical        Vec3
	LowerLeftCorner Vec3
	LensRadius      float32
	U, V, W         Vec3
}

// NewCamera gets a camera with some defaults
func NewCamera(lookFrom, lookAt, up Vec3, fov, aspect, aperture, focusDistance float32) Camera {
	// var u, v, w Vec3

	var camera Camera

	camera.LensRadius = aperture / 2

	theta := fov * math32.Pi / 180
	halfHeight := math32.Tan(theta / 2)
	halfWidth := aspect * halfHeight

	camera.Origin = lookFrom

	camera.W = lookFrom.Subtract(lookAt).Normalized()
	camera.U = up.Cross(camera.W).Normalized()
	camera.V = camera.W.Cross(camera.U)

	// camera.LowerLeftCorner = Vec3{-halfWidth, -halfHeight, -1.0}
	camera.LowerLeftCorner = camera.Origin.
		Subtract(camera.U.ScalarMultiply(halfWidth * focusDistance)).
		Subtract(camera.V.ScalarMultiply(halfHeight * focusDistance)).
		Subtract(camera.W.ScalarMultiply(focusDistance))
	camera.Horizontal = camera.U.ScalarMultiply(2.0 * halfWidth * focusDistance)
	camera.Vertical = camera.V.ScalarMultiply(2.0 * halfHeight * focusDistance)

	return camera
}

// GetRay gets the ray from the camera.
func (c Camera) GetRay(s, t float32) Ray {
	rd := RandomInUnitSphere().ScalarMultiply(c.LensRadius)
	offset := c.U.ScalarMultiply(rd.X()).Add(c.V.ScalarMultiply(rd.Y()))
	return Ray{
		c.Origin.Add(offset),
		c.LowerLeftCorner.
			Add(c.Horizontal.ScalarMultiply(s)).
			Add(c.Vertical.ScalarMultiply(t)).
			Subtract(c.Origin).
			Subtract(offset),
	}
}

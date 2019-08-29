package main

import "github.com/chewxy/math32"

// Refract performs a refraction.
func Refract(v Vec3, n Vec3, niOverNt float32, refracted *Vec3) bool {
	uv := v.Normalized()
	dt := uv.Dot(n)
	discriminant := 1.0 - niOverNt*niOverNt*(1-(dt*dt))
	if discriminant > 0 {
		refracted.Assign(
			uv.
				Subtract(n.ScalarMultiply(dt)).
				ScalarMultiply(niOverNt).
				Subtract(
					n.ScalarMultiply(math32.Sqrt(discriminant)),
				),
		)
		return true
	}
	return false
}

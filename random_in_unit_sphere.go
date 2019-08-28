package main

import "math/rand"

// RandomInUnitSphere gets a random vector in a unit sphere.
func RandomInUnitSphere() Vec3 {
	for {
		randVec := Vec3{rand.Float32(), rand.Float32(), rand.Float32()}
		inCube := randVec.Subtract(Vec3{0.5, 0.5, 0.5}).ScalarMultiply(2.0)
		if inCube.Length() <= 1.0 {
			return inCube
		}
	}
}

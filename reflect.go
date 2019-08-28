package main

// Reflect returns a reflected vector
func Reflect(v Vec3, n Vec3) Vec3 {
	return v.Subtract(n.ScalarMultiply(2 * v.Dot(n)))
}

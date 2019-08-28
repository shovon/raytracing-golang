package main

import (
	"fmt"

	"github.com/chewxy/math32"
)

// Vec3 represents a 3D vector
type Vec3 struct {
	e0 float32
	e1 float32
	e2 float32
}

// X represents first component of the vector
func (v Vec3) X() float32 {
	return v.e0
}

// Y represents second component of the vector
func (v Vec3) Y() float32 {
	return v.e1
}

// Z represents third component of the vector
func (v Vec3) Z() float32 {
	return v.e2
}

// R represents red (first) component of the vector
func (v Vec3) R() float32 {
	return v.e0
}

// G represents green (second) component of the vector
func (v Vec3) G() float32 {
	return v.e1
}

// B represents blue (third) component of the vector
func (v Vec3) B() float32 {
	return v.e2
}

func (v Vec3) String() string {
	return fmt.Sprintf("%f %f %f", v.e0, v.e1, v.e2)
}

// Add adds each of the two vectors' components, and returns a new vector with
// the components summed
func (v Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{v.e0 + v2.e0, v.e1 + v2.e1, v.e2 + v2.e2}
}

// Subtract subtractss each of the two vectors' components, and returns a new
// vector with the components summed.
func (v Vec3) Subtract(v2 Vec3) Vec3 {
	return Vec3{v.e0 - v2.e0, v.e1 - v2.e1, v.e2 - v2.e2}
}

// Hadamard multiplies each of the two vectors' components, and returns a new
// vector with the components multiplied. (Alias for multiply.)
func (v Vec3) Hadamard(v2 Vec3) Vec3 {
	return Vec3{v.e0 * v2.e0, v.e1 * v2.e1, v.e2 * v2.e2}
}

// Multiply multiplies each of the two vectors' components, and returns a new
// vector with the components multiplied. (Alias for Hadamard)
func (v Vec3) Multiply(v2 Vec3) Vec3 {
	return v.Hadamard(v2)
}

// Divide divides each of the two vectors' components, and returns a new vector
// with the components divided.
func (v Vec3) Divide(v2 Vec3) Vec3 {
	return Vec3{v.e0 / v2.e0, v.e1 / v2.e1, v.e2 / v2.e2}
}

// ScalarMultiply multiplies each of the component of the vector by the supplied
// floating point value.
func (v Vec3) ScalarMultiply(f float32) Vec3 {
	return Vec3{f, f, f}.Multiply(v)
}

// ScalarDivide divides each of the component of the vector by the supplied
// floating point value.
func (v Vec3) ScalarDivide(f float32) Vec3 {
	return v.Multiply(Vec3{1 / f, 1 / f, 1 / f})
}

// Dot performs the dot product between the two vectors.
func (v Vec3) Dot(v2 Vec3) float32 {
	return v.e0*v2.e0 + v.e1*v2.e1 + v.e2*v2.e2
}

// Cross performs the cross product
func (v Vec3) Cross(v2 Vec3) Vec3 {
	return Vec3{
		v.e1*v2.e2 - v.e2*v2.e1,
		-(v.e0*v2.e2 - v.e2*v.e0),
		v.e0*v.e2 - v.e1*v2.e0,
	}
}

// Length gets the distance of the vector from the origin.
func (v Vec3) Length() float32 {
	return math32.Sqrt(v.Dot(v))
}

// Normalized gets the normalized form of the vector.
func (v Vec3) Normalized() Vec3 {
	return v.ScalarDivide(v.Length())
}

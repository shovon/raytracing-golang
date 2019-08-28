package main

// Material represents the material of the given object.
type Material interface {
	Scatter(rIn *Ray, rec *HitRecord, attenuation *Vec3, scattered *Ray) bool
}

package main

type Lambertian struct {
	Albedo Vec3
}

func (l Lambertian) Scatter(
	rIn *Ray,
	rec *HitRecord,
	attenuation *Vec3,
	scattered *Ray,
) bool {
	return false
}

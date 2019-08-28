package main

// Lambertian this is the lambertian material.
type Lambertian struct {
	Albedo Vec3
}

// Scatter scatter the incoming ray.
func (l Lambertian) Scatter(
	rIn *Ray,
	rec *HitRecord,
	attenuation *Vec3,
	scattered *Ray,
) bool {
	target := rec.P.Add(rec.Normal).Add(RandomInUnitSphere())
	scattered.Assign(Ray{rec.P, target.Subtract(rec.P)})
	attenuation.Assign(l.Albedo)
	return true
}

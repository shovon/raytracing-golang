package main

import "math/rand"

// Dialectric is a see-through dialectric material.
type Dialectric struct {
	RefractionIndex float32
}

// Scatter logic for scattering the ray when it hits the dialectric material.
func (d Dialectric) Scatter(
	rIn *Ray,
	rec *HitRecord,
	attenuation *Vec3,
	scattered *Ray,
) bool {
	outwardNormal := Vec3{}
	reflected := Reflect(rIn.Direction(), rec.Normal)
	var niOverNt float32
	attenuation.Assign(Vec3{1.0, 1.0, 1.0})
	var refracted Vec3
	var reflectProb float32
	var cosine float32
	if rIn.Direction().Dot(rec.Normal) > 0.0 {
		outwardNormal = rec.Normal.ScalarMultiply(-1.0)
		niOverNt = d.RefractionIndex
		cosine = d.RefractionIndex * rIn.Direction().Dot(rec.Normal) / rIn.Direction().Length()
	} else {
		outwardNormal = rec.Normal
		niOverNt = 1.0 / d.RefractionIndex
		cosine = -rIn.Direction().Dot(rec.Normal) / rIn.Direction().Length()
	}
	if Refract(rIn.Direction(), outwardNormal, niOverNt, &refracted) {
		reflectProb = Schlick(cosine, d.RefractionIndex)
	} else {
		scattered.Assign(Ray{rec.P, reflected})
		reflectProb = 1
	}
	if rand.Float32() < reflectProb {
		scattered.Assign(Ray{rec.P, reflected})
	} else {
		scattered.Assign(Ray{rec.P, refracted})
	}
	return true
}

package main

type Metal struct {
	Albedo Vec3
	Fuzz   float32
}

func (m Metal) Scatter(
	rIn *Ray,
	rec *HitRecord,
	attenuation *Vec3,
	scattered *Ray,
) bool {
	f := m.Fuzz
	// if f > 1.0 {
	// 	f = 1.0
	// }
	reflected := Reflect(rIn.Direction().Normalized(), rec.Normal)
	scattered.Assign(Ray{rec.P, reflected.Add(RandomInUnitSphere().ScalarMultiply(f))})
	attenuation.Assign(m.Albedo)
	return scattered.Direction().Dot(rec.Normal) > 0
}

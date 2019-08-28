package main

type Metal struct {
	Albedo Vec3
}

func (m Metal) Scatter(
	rIn *Ray,
	rec *HitRecord,
	attenuation *Vec3,
	scattered *Ray,
) bool {
	reflected := Reflect(rIn.Direction().Normalized(), rec.Normal)
	scattered.Assign(Ray{rec.P, reflected})
	attenuation.Assign(m.Albedo)
	return scattered.Direction().Dot(rec.Normal) > 0
}

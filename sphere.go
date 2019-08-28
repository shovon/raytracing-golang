package main

import "github.com/chewxy/math32"

// Sphere represents a hitable object that is a sphere
type Sphere struct {
	Center   Vec3
	Radius   float32
	Material Material
}

// Hit determines whether or not the given ray was hit.
func (s Sphere) Hit(r Ray, tMin float32, tMax float32, rec *HitRecord) bool {
	oc := r.Origin().Subtract(s.Center)
	a := r.Direction().Dot(r.Direction())
	b := oc.Dot(r.Direction())
	c := oc.Dot(oc) - s.Radius*s.Radius
	discriminant := b*b - a*c
	if discriminant > 0 {
		temp := (-b - math32.Sqrt(b*b-a*c)) / a
		if temp < tMax && temp > tMin {
			rec.T = temp
			rec.P = r.PointAtParameter(rec.T)
			rec.Normal = rec.P.Subtract(s.Center).ScalarDivide(s.Radius)
			rec.Material = s.Material
			return true
		}
		temp = (-b + math32.Sqrt(b*b-a*c)) / a
		if temp < tMax && temp > tMin {
			rec.T = temp
			rec.P = r.PointAtParameter(rec.T)
			rec.Normal = rec.P.Subtract(s.Center).ScalarDivide(s.Radius)
			rec.Material = s.Material
			return true
		}
	}
	return false
}

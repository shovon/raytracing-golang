package main

// Hitable a hitable type
type Hitable interface {
	Hit(r Ray, tMin float32, tMax float32, rec *HitRecord) bool
}

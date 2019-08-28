package main

func (h *HitRecord) Assign(n HitRecord) {
	h.T = n.T
	h.P = n.P
	h.Normal = n.Normal
}

// Hitable a hitable type
type Hitable interface {
	Hit(r Ray, tMin float32, tMax float32, rec *HitRecord) bool
}

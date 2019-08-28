package main

// HitRecord a hit record
type HitRecord struct {
	T        float32
	P        Vec3
	Normal   Vec3
	Material Material
}

// Assign assigns the values to this instance of HitRecord
func (h *HitRecord) Assign(n HitRecord) {
	h.T = n.T
	h.P = n.P
	h.Normal = n.Normal
	h.Material = n.Material
}

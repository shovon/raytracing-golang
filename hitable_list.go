package main

// HitableList a list of hitable objects (not sure what this is for).
type HitableList struct {
	List []Hitable
}

// Hit determines if hit.
func (h HitableList) Hit(
	r Ray,
	tMin float32,
	tMax float32,
	rec *HitRecord,
) bool {
	var tempRec HitRecord
	hitAnything := false
	closestSoFar := tMax
	for _, hitable := range h.List {
		if hitable.Hit(r, tMin, closestSoFar, &tempRec) {
			hitAnything = true
			closestSoFar = tempRec.T
			rec.Assign(tempRec)
		}
	}
	return hitAnything
}

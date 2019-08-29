package main

import "github.com/chewxy/math32"

// Schlick this gives glass reflexivity
func Schlick(cosine, refractionIndex float32) float32 {
	r0 := (1 - refractionIndex) / (1 + refractionIndex)
	r0 = r0 * r0
	return r0 + (1-r0)*math32.Pow((1-cosine), 5)
}

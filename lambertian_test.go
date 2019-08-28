package main

import (
	"testing"
)

func TestLambertian(t *testing.T) {
	var m Material
	m = &Lambertian{}
	if m == nil {
		t.Error("Should not be here")
	}
}

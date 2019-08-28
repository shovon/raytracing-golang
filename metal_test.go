package main

import "testing"

func TestMetal(t *testing.T) {
	var m Material
	m = &Metal{}
	if m == nil {
		t.Error("Should not be here")
	}
}

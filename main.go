package main

import (
	"fmt"
	"os"

	"github.com/chewxy/math32"
)

func didHitSphere(center Vec3, radius float32, r Ray) bool {
	oc := r.Origin().Subtract(center)
	a := r.Direction().Dot(r.Direction())
	b := 2.0 * oc.Dot(r.Direction())
	c := oc.Dot(oc) - radius*radius
	discriminant := b*b - 4*a*c
	return discriminant > 0
}

func hitSphere(center Vec3, radius float32, r Ray) float32 {
	oc := r.Origin().Subtract(center)
	a := r.Direction().Dot(r.Direction())
	b := 2.0 * oc.Dot(r.Direction())
	c := oc.Dot(oc) - radius*radius
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return -1.0
	}
	return (-b - math32.Sqrt(discriminant)) / (2.0 * a)
}

func drawPixel(col Vec3) {
	ir := int(255.99 * col.R())
	ig := int(255.99 * col.G())
	ib := int(255.99 * col.B())
	fmt.Fprintf(os.Stdout, "%d %d %d\n", ir, ig, ib)
}

func main() {
	nx, ny := 200, 100

	fmt.Fprintf(os.Stdout, "P3\n%d %d\n255\n", nx, ny)

	lowerLeftCorner := Vec3{-2.0, -1.0, -1.0}
	horizontal := Vec3{4.0, 0.0, 0.0}
	vertical := Vec3{0.0, 2.0, 0.0}
	origin := Vec3{0.0, 0.0, 0.0}

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float32(i) / float32(nx)
			v := float32(j) / float32(ny)
			r := Ray{
				origin,
				lowerLeftCorner.
					Add(horizontal.ScalarMultiply(u)).
					Add(vertical.ScalarMultiply(v)),
			}
			col := color(r)
			drawPixel(col)
		}
	}
}

func color(r Ray) Vec3 {
	// This is the old code for just a boolean sphere hitting logic.
	// if didHitSphere(Vec3{0.0, 0.0, -1.0}, 0.5, r) {
	// 	return Vec3{1.0, 0.0, 0.0}
	// }
	t := hitSphere(Vec3{0.0, 0.0, -1.0}, 0.5, r)
	if t > 0.0 {
		n := r.PointAtParameter(t).Subtract(Vec3{0.0, 0.0, -1.0})
		return Vec3{n.X() + 1, n.Y() + 1, n.Z() + 1}.ScalarMultiply(0.5)
	}
	unitDirection := r.Direction().Normalized()
	t = 0.5*unitDirection.Y() + 1.0
	return Vec3{1.0, 1.0, 1.0}.ScalarMultiply(1.0 - t).
		Add(Vec3{0.5, 0.7, 1.0}.ScalarMultiply(t))
}

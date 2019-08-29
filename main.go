package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

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

func randomScene() []Hitable {
	list := make([]Hitable, 0, 10)
	list = append(list, Sphere{Vec3{0, -1000, 0}, 1000, Lambertian{Vec3{0.5, 0.5, 0.5}}})
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMaterial := rand.Float32()
			center := Vec3{float32(a) + (0.9 * rand.Float32()), 0.2, float32(b) + (0.9 * rand.Float32())}
			if center.Add(Vec3{0, 0.8, 0}).Subtract(Vec3{0, 1, 0}).Length() <= 1 {
				continue
			}
			if center.Subtract(Vec3{4, 0.2, 0}).Length() > 0.9 {
				if chooseMaterial < 0.8 {
					list = append(list, Sphere{center, 0.2, Lambertian{Vec3{rand.Float32(), rand.Float32(), rand.Float32()}}})
				} else if chooseMaterial < 0.95 {
					list = append(
						list,
						Sphere{
							center,
							0.2,
							Metal{
								Vec3{
									0.5 * (1 + rand.Float32()),
									0.5 * (1 + rand.Float32()),
									0.5 * (1 + rand.Float32()),
								},
								0.5 * (1 + rand.Float32()),
							},
						},
					)
				} else {
					list = append(list, Sphere{center, 0.2, Dialectric{1.5}})
				}
			}
		}
	}

	list = append(list, Sphere{Vec3{0, 1, 0}, 1.0, Dialectric{1.5}})
	list = append(list, Sphere{Vec3{-4, 1, 0}, 1.0, Lambertian{Vec3{0.4, 0.2, 0.1}}})
	list = append(list, Sphere{Vec3{4, 1, 0}, 1.0, Metal{Vec3{0.7, 0.6, 0.5}, 0.0}})

	return list
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// nx, ny, ns := 100, 50, 100
	nx, ny, ns := 1024, 576, 100

	fmt.Fprintf(os.Stdout, "P3\n%d %d\n255\n", nx, ny)

	// r := math32.Cos(math32.Pi / 4)

	// hitable := []Hitable{
	// 	Sphere{Vec3{0.0, 0.0, -1.0}, 0.5, Lambertian{Vec3{0.8, 0.3, 0.3}}},
	// 	Sphere{Vec3{0.0, -100.5, -1.0}, 100.0, Lambertian{Vec3{0.8, 0.8, 0.0}}},
	// 	Sphere{Vec3{1.0, 0.0, -1.0}, 0.5, Metal{Vec3{0.8, 0.6, 0.2}, 0.01}},
	// 	Sphere{Vec3{-1.0, 0.0, -1.0}, 0.5, Dialectric{1.5}},

	// 	// Sphere{Vec3{-r, 0.0, -1.0}, r, Lambertian{Vec3{0.0, 0.0, 1.0}}},
	// 	// Sphere{Vec3{r, 0.0, -1}, r, Lambertian{Vec3{1.0, 0.0, 0.0}}},
	// }

	hitable := randomScene()

	world := HitableList{hitable}

	lookFrom := Vec3{12, 2, 4}
	lookAt := Vec3{0, 1, 0}
	distToFocus := (lookFrom.Subtract(lookAt)).Length()
	aparture := float32(0.2)
	cam := NewCamera(lookFrom, lookAt, Vec3{0, 1, 0}, 20, float32(nx)/float32(ny), aparture, distToFocus)

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			col := Vec3{0.0, 0.0, 0.0}
			for s := 0; s < ns; s++ {
				u := (float32(i) + rand.Float32()) / float32(nx)
				v := (float32(j) + rand.Float32()) / float32(ny)
				r := cam.GetRay(u, v)
				// p := r.PointAtParameter(2.0) // ?
				col = col.Add(color(r, world, 0))
			}
			col = col.ScalarDivide(float32(ns))
			drawPixel(Vec3{
				math32.Sqrt(col.R()),
				math32.Sqrt(col.G()),
				math32.Sqrt(col.B()),
			})
		}
	}
}

func color(r Ray, world Hitable, depth int) Vec3 {
	var rec HitRecord
	if world.Hit(r, 0.001, math.MaxFloat32, &rec) {
		var scattered Ray
		var attenuation Vec3
		if depth < 50 && rec.Material.Scatter(&r, &rec, &attenuation, &scattered) {
			return attenuation.Hadamard(color(scattered, world, depth+1))
		}
		return Vec3{0.0, 0.0, 0.0}
		// target := rec.P.Add(rec.Normal).Add(RandomInUnitSphere())
		// return color(Ray{rec.P, target.Subtract(rec.P)}, world).ScalarMultiply(0.5)
	}

	unitDirection := r.Direction().Normalized()
	t := 0.5*unitDirection.Y() + 1.0
	return Vec3{1.0, 1.0, 1.0}.ScalarMultiply(1.0 - t).
		Add(
			Vec3{0.5, 0.7, 1.0}.
				ScalarMultiply(t),
		)
}

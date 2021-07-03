package main

import (
	"fmt"
	"math"

	"github.com/fogleman/gg"
)

type Point struct {
	x, y, z float64
}

type Color struct {
	r, g, b int
}

type Sphere struct {
	center Point
	radius float64
	color  Color
}

var dc gg.Context // graphics context

var spheres [3]Sphere          // spheres in scene
var Cw, Ch int                 // canvas
var Vw, Vh float64             // viewport
var projection_plane_d float64 // projection plane
var origin Point               // origin
var BACKGROUND_COLOR Color

// ok
func CanvasToViewport(x, y int) Vector {
	var newX, newY float64
	newX = float64(x) * Vw / float64(Cw)
	newY = float64(y) * Vh / float64(Ch)
	return Vector{newX, newY, projection_plane_d}
}

// ok
func PutPixel(x, y int, color Color) {
	newX := Cw/2 + x
	newY := Ch/2 - y - 1
	dc.SetRGB255(color.r, color.g, color.b)
	dc.SetPixel(newX, newY)
}

// ok
func IntersectRaySphere(O Point, D Vector, sphere Sphere) (float64, float64) {
	r := sphere.radius
	var CO Vector
	CO.X = O.x - sphere.center.x
	CO.Y = O.y - sphere.center.y
	CO.Z = O.z - sphere.center.z

	a := D.Dot(D)
	b := 2 * CO.Dot(D)
	c := CO.Dot(CO) - r*r

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return math.Inf(1), math.Inf(1)
	}

	t1 := (-b + math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b - math.Sqrt(discriminant)) / (2 * a)
	return t1, t2
}

func TraceRay(O Point, D Vector, t_min, t_max float64) Color {
	closest_t := math.Inf(1)
	closest_sphere := -1
	for index, sphere := range spheres {
		t1, t2 := IntersectRaySphere(O, D, sphere)
		if t1 > t_min && t1 < t_max && t1 < closest_t {
			closest_t = t1
			closest_sphere = index
		}
		if t2 > t_min && t2 < t_max && t2 < closest_t {
			closest_t = t2
			closest_sphere = index
		}
	}
	if closest_sphere == -1 {
		return BACKGROUND_COLOR
	}

	return spheres[closest_sphere].color
}

func main() {
	Cw, Ch = 1024, 1024
	Vw, Vh = 1, 1 // 1 x 1
	projection_plane_d = 1
	origin = Point{0, 0, 0}
	BACKGROUND_COLOR = Color{255, 255, 255}

	// spheres
	spheres[0].center = Point{0, -1, 3}
	spheres[0].radius = 1
	spheres[0].color = Color{255, 0, 0} // red

	spheres[1].center = Point{2, 0, 4}
	spheres[1].radius = 1
	spheres[1].color = Color{0, 0, 255} // blue

	spheres[2].center = Point{-2, 0, 4}
	spheres[2].radius = 1
	spheres[2].color = Color{0, 255, 0} // green

	// create context
	dc = *gg.NewContext(Cw, Ch)

	// render image
	fmt.Println("rendering image")
	for x := -Cw / 2.0; x < Cw/2.0; x++ {
		for y := -Ch / 2.0; y < Ch/2.0; y++ {
			D := CanvasToViewport(x, y)
			color := TraceRay(origin, D, 1, math.Inf(1))
			PutPixel(int(x), int(y), color)
		}

	}

	// save it
	fmt.Println("saving image")
	dc.SavePNG("out.png")
}

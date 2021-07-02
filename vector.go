package main

// https://www.netguru.com/blog/vector-operations-in-go
import (
	"fmt"
	"math"
)

// Vector - struct holding X Y Z values of a 3D vector
type Vector struct {
	X, Y, Z float64
}

func (a Vector) Add(b Vector) Vector {
	return Vector{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

func (a Vector) Sub(b Vector) Vector {
	return Vector{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

func (a Vector) MultiplyByScalar(s float64) Vector {
	return Vector{
		X: a.X * s,
		Y: a.Y * s,
		Z: a.Z * s,
	}
}

func (a Vector) Dot(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a Vector) Length() float64 {
	return math.Sqrt(a.Dot(a))
}

func (a Vector) Cross(b Vector) Vector {
	return Vector{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

func (a Vector) Normalize() Vector {
	return a.MultiplyByScalar(1. / a.Length())
}

func TestVector() {
	result := Vector{1., 1., 1.}.Add(Vector{2., 2., 2.})
	fmt.Println(result)

	result = Vector{3., 3., 3.}.Sub(Vector{1., 1., 1.})
	fmt.Println(result)

	scalar := Vector{1., 0., 0.}.Dot(Vector{0., 1., 0.})
	fmt.Println("perpendicular: ", scalar)

	scalar = Vector{1., 0., 0.}.Dot(Vector{1., 0., 0.})
	fmt.Println("parallel: ", scalar)

	scalar = Vector{1, 2, 2}.Length()
	fmt.Println("length: ", scalar)

	result = Vector{1., 0., 0.}.Cross(Vector{0., 1., 0.})
	fmt.Println("cross product is perpendicular: ", result)

	result = Vector{1., 0., 0.}.Cross(Vector{0., 0., 1.})
	fmt.Println("cross product is perpendicular: ", result)
}

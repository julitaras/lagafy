// Package triangle does interesting stuff with triangles.
package triangle

const (
	testVersion = 3

	NaT = 0 // not a triangle
	Equ = 1 // equilateral
	Iso = 2 // isosceles
	Sca = 3 // scalene
)

type Kind int8

// KindFromSides returns the Kind of a triangle.
func KindFromSides(a, b, c float64) Kind {
	return 0
}

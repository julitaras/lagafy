// Package triangle does interesting stuff with triangles.
package triangle

import "math"

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
	if invalid(a) || invalid(b) || invalid(c) {
		return NaT
	}

	if (a+b) < c || (b+c) < a || (c+a) < b {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || b == c || c == a {
		return Iso
	}

	return Sca
}

func invalid(x float64) bool {
	return x == 0 || math.IsNaN(x) || math.IsInf(x, 0)
}

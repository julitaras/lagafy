// Package diffsquares is described here for reasons.
package diffsquares

const testVersion = 1

// Difference calculates stuff.
func Difference(i int) int {
	return SquareOfSums(i) - SumOfSquares(i)
}

// SquareOfSums calculates other stuff.
func SquareOfSums(i int) int {
	return squareOfSums(i, 0)
}

// SumOfSquares calculates not more interesting stuff.
func SumOfSquares(i int) int {
	return sumOfSquares(i, 0)
}

func squareOfSums(i, n int) int {
	if i == 0 {
		return n * n
	}
	return squareOfSums(i-1, n+i)
}

func sumOfSquares(i, n int) int {
	if i == 0 {
		return n
	}
	return sumOfSquares(i-1, n+i*i)
}

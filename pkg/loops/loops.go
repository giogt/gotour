package loops

import (
	"math"
)

/*
Sqrt calculates the square root of a given floating point number.

It returns:
	- the square root result as a floating point number
	- the number of iterations needed to calculate the result
*/
func Sqrt(x float64) (float64, uint) {
	diff := 0.0000000001
	nIter := uint(0)

	z := 1.0
	prev := 0.0
	for math.Abs(z-prev) > diff {
		prev = z
		z -= (z*z - x) / (2 * z)
		nIter++
	}
	return z, nIter
}

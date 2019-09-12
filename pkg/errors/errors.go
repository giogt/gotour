package errors

import (
	"fmt"
	"math"
)

/*
Sqrt calculates the square root of a given floating point number.

It returns:
	- the square root result as a floating point number
	- the number of iterations needed to calculate the result
	- a ErrNegativeSqrt error, when trying to calculate the square root of
	  a negative number
*/
func Sqrt(x float64) (float64, uint, error) {
	if x < 0.0 {
		return 0, 0, ErrNegativeSqrt(x)
	}

	diff := 0.0000000001
	nIter := uint(0)

	z := 1.0
	prev := 0.0
	for math.Abs(z-prev) > diff {
		prev = z
		z -= (z*z - x) / (2 * z)
		nIter++
	}
	return z, nIter, nil
}

/*
ErrNegativeSqrt describes an error occurring when trying to compute the square
root of a negative number
*/
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

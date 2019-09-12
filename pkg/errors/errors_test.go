package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Sqrt_forAPerfectSquare_mustReturnTheExactResult(t *testing.T) {
	res, _, err := Sqrt(25)
	assert.Equal(t, float64(5), res)
	assert.Nil(t, err)
}

func Test_Sqrt_forNumberTwo_mustReturnAGoodApproximation(t *testing.T) {
	res, _, err := Sqrt(2)

	lowerBound := float64(1.4142135623730950)
	higherBound := float64(1.4142135623730951)
	assert.Condition(t, func() bool { return res >= lowerBound && res <= higherBound })
	assert.Nil(t, err)
}

func Test_Sqrt_forNegativeNumber_mustReturnAnError(t *testing.T) {
	res, _, err := Sqrt(-1)
	assert.Equal(t, float64(0), res)
	assert.Equal(t, "cannot Sqrt negative number: -1", err.Error())
}

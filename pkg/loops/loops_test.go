package loops

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Sqrt_forAPerfectSquare_mustReturnTheExactResult(t *testing.T) {
	res, _ := Sqrt(25)
	assert.Equal(t, float64(5), res)
}

func Test_Sqrt_forNumberTwo_mustReturnAGoodApproximation(t *testing.T) {
	res, _ := Sqrt(2)

	lowerBound := float64(1.4142135623730950)
	higherBound := float64(1.4142135623730951)
	assert.Condition(t, func() bool { return res >= lowerBound && res <= higherBound })
}

func Test_Sqrt_forNegativeNumber_mustReturnNan(t *testing.T) {
	res, _ := Sqrt(-1)
	assert.True(t, math.IsNaN(res))
}

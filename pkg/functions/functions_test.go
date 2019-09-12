package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nextFibonacci(t *testing.T) {
	const n = 10

	nextF := nextFibonacci()
	var f [n]uint64
	for i := 0; i < n; i++ {
		f[i] = nextF()
	}
	assert.Equal(t, [n]uint64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}, f)
}

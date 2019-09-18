package channels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fibonacci(t *testing.T) {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	var res []int
	for n := range c {
		res = append(res, n)
	}

	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	assert.Equal(t, expected, res)
}

package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Walk(t *testing.T) {
	var ch = make(chan int)

	go Walk(New(1), ch)
	var res []int
	for v := range ch {
		res = append(res, v)
	}

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, expected, res)
}

func Test_Same(t *testing.T) {
	assert.True(t, Same(New(1), New(1)))
	assert.False(t, Same(New(1), New(2)))
}

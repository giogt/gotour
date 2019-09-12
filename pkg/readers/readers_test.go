package readers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_InfiniteCharStreamReader(t *testing.T) {
	// let's test the first 8 * 4096 chars
	const size = 8
	for i := 0; i < 4096; i++ {
		b := make([]byte, size)
		n, err := InfiniteCharStreamReader{}.Read(b)
		assert.Equal(t, size, n)
		assert.Equal(t, b, []byte{'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A'})
		assert.Nil(t, err)
	}
}

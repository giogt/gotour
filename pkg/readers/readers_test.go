package readers

import (
	"strings"
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

func Test_rot13Reader(t *testing.T) {
	const s = "Lbh penpxrq gur pbqr!"
	const expected = "You cracked the code!"

	reader := rot13Reader{strings.NewReader(s)}
	var res []byte = make([]byte, len(s))
	reader.Read(res)

	assert.Equal(t, expected, string(res))
}

func Test_rot13(t *testing.T) {
	const s = "The Quick Brown Fox Jumps Over The Lazy Dog"
	const expected = "Gur Dhvpx Oebja Sbk Whzcf Bire Gur Ynml Qbt"

	var res []byte = make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = rot13(s[i])
	}

	assert.Equal(t, expected, string(res))
}

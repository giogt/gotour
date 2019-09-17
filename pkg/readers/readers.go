package readers

import "io"

/*
InfiniteCharStreamReader is a reader that emits an infinite stream of
the ASCII character 'A'.
*/
type InfiniteCharStreamReader struct{}

func (r InfiniteCharStreamReader) Read(b []byte) (n int, err error) {
	n = len(b)
	err = nil // this infinite stream reader will never fail and always produce characters

	for i := 0; i < n; i++ {
		b[i] = 'A'
	}
	return
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return
}

func rot13(c byte) byte {
	const offset = 13
	const azLength = 26

	var res byte

	if c >= 'a' && c <= 'z' {
		res = c + offset
		if res > 'z' {
			res -= azLength
		}
	} else if c >= 'A' && c <= 'Z' {
		res = c + offset
		if res > 'Z' {
			res -= azLength
		}
	} else {
		res = c
	}

	return res
}

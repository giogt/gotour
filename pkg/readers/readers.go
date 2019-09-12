package readers

/*
InfiniteCharStreamReader is a reader that emits an infinite stream of
the ASCII character 'A'.
*/
type InfiniteCharStreamReader struct{}

func (r InfiniteCharStreamReader) Read(b []byte) (n int, err error) {
	cb := len(b)
	for i := 0; i < cb; i++ {
		b[i] = 'A'
	}
	return cb, nil
}

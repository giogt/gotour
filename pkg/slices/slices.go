package slices

func Pic(dx, dy int) [][]uint8 {
	// allocate matrix
	p := make([][]uint8, dy)
	for i, _ := range p {
		p[i] = make([]uint8, dx)
	}

	for i := range p {
		for j := range p[i] {
			//p[i][j] = uint8((i + j) / 2)
			//p[i][j] = uint8(i * j)
			p[i][j] = uint8(i ^ j)
		}
	}

	return p
}

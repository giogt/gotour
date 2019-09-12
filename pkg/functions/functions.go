package functions

func nextFibonacci() func() uint64 {
	n := uint64(0)
	prev1 := uint64(0)
	prev2 := uint64(0)
	return func() uint64 {
		var res uint64

		// increment n after every execution
		defer func() { n++ }()

		switch n {
		case 0: res = 0
		case 1: res = 1
		default: res = prev1 + prev2
		}

		prev2 = prev1
		prev1 = res
		return res
	}
}

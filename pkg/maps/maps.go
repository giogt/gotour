package maps

import (
	"fmt"
	"strings"
)

func MutateMap() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func WordCount(s string) map[string]int {
	wc := make(map[string]int)
	for _, w := range strings.Fields(s) {
		wc[w]++
	}
	return wc
}

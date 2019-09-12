package maps

import (
	"testing"

	"golang.org/x/tour/wc"
)

func Test_wordCount(t *testing.T) {
	wc.Test(WordCount)
}

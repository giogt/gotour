package crawler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Crawl(t *testing.T) {
	res := Crawl("https://golang.org/", 4, fetcher)
	for k, v := range fetcher {
		r, ok := res[k]
		assert.True(t, ok)
		assert.Equal(t, v.body, r.body)
		assert.Equal(t, v.urls, r.urls)
	}
}

// mockFetcher is Fetcher that returns mock results, used for testing
type mockFetcher map[string]*mockResult

type mockResult struct {
	body string
	urls []string
}

func (f mockFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated mockFetcher
var fetcher = mockFetcher{
	"https://golang.org/": &mockResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &mockResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &mockResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &mockResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

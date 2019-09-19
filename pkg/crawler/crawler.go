package crawler

import (
	"errors"
	"sync"
)

// FetchResult describe the result of a fetch for a given URL
type FetchResult struct {
	body string
	urls []string
	err  error
}

/*
A Fetcher returns the body of URL and a slice of URLs found on that page.
*/
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

var fetchersWg = sync.WaitGroup{}

/*
Crawl uses fetcher to recursively crawl pages starting with url, to a maximum of depth.
It returns a map where:
	- the key is a URL
	- the value is the fetch result for that URL, consisting of:
		- the page body for that URL
		- the URLs contained in the page
		- an error, valued if the fetching for the URL failed
*/
func Crawl(url string, depth int, fetcher Fetcher) map[string]FetchResult {
	res := make(map[string]FetchResult)

	fetchersWg.Add(1)
	go crawl(url, depth, fetcher, res)
	fetchersWg.Wait()

	return res
}

func crawl(url string, depth int, fetcher Fetcher, res map[string]FetchResult) {
	defer fetchersWg.Done()

	if depth <= 0 {
		return
	}

	exists := putIfAbsent(res, url, FetchResult{"", nil, errors.New("fetching URL in progress")})
	if exists {
		// another fetcher is already downloading / has already downloaded the URL
		return
	}

	body, urls, err := fetcher.Fetch(url)
	res[url] = FetchResult{body, urls, err}
	if err != nil {
		return
	}

	for _, u := range urls {
		fetchersWg.Add(1)
		go crawl(u, depth-1, fetcher, res)
	}
}

var mapMux = sync.Mutex{}

func putIfAbsent(m map[string]FetchResult, url string, f FetchResult) bool {
	defer mapMux.Unlock()
	mapMux.Lock()

	_, exists := m[url]
	if !exists {
		m[url] = f
	}
	return exists
}

package main

import (
	"fmt"
	"sync"
)

// SafeMap is safe to use concurrently.
type SafeMap struct {
	v   map[string]bool
	mux sync.Mutex
}

// SetVal sets the value for the given key.
func (m *SafeMap) SetVal(key string, val bool) {
	m.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	m.v[key] = val
	m.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (m *SafeMap) GetVal(key string) bool {
	m.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer m.mux.Unlock()
	return m.v[key]
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, status chan bool, urlMap *SafeMap) {

	// Check if we fetched this url previously.
	if ok := urlMap.GetVal(url); ok {
		//fmt.Println("Already fetched url!")
		status <- true
		return
	}

	// Marking this url as fetched already.
	urlMap.SetVal(url, true)

	if depth <= 0 {
		status <- false
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		status <- false
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	statuses := make([]chan bool, len(urls))
	for index, u := range urls {
		statuses[index] = make(chan bool)
		go Crawl(u, depth-1, fetcher, statuses[index], urlMap)
	}

	// Wait for child goroutines.
	for _, childstatus := range statuses {
		<-childstatus
	}

	// And now this goroutine can finish.
	status <- true

	return
}

func main() {
	urlMap := SafeMap{v: make(map[string]bool)}
	status := make(chan bool)
	go Crawl("https://golang.org/", 4, fetcher, status, &urlMap)
	<-status
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

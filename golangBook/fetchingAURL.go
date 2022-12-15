// Fetch prints the content found at a URL.
package main

/**
 * Run command: `go run . http://gopl.io`
 */

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func fetchUrl() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

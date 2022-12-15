// Server1 is a minimal "echo" server.
package main

/**
 * Run command: `go run . https://golang.org`
 * Open web browser: http://localhost:8000
 * Web sẽ trả về HTML của trang https://golang.org
 */

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func webServerfetchUrl(w http.ResponseWriter, r *http.Request) {
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
		fmt.Fprintf(w, "%s", b)
	}
}

func webServer() {
	http.HandleFunc("/", webServerfetchUrl) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

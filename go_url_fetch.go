//fetch url from specified site, prints contents
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	//timer
	start := time.Now()
	//fetch url
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//reading url data. note the read is closed once entire doc is read
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		//copy and read url data
		doc, err := io.Copy(ioutil.Discard, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", b, err)

		}
		fmt.Printf("%s", doc)
		fmt.Printf("%.2fs elasped\n", time.Since(start).Seconds())

	}

}

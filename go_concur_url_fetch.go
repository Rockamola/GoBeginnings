//fetch urls concurrently. main function runs in a goroutine and
//the go statement creates additional goroutines
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
	//receive-channel
	ch := make(chan string)
	//error channel/print output
	errCh := make(chan string)
	//done channel/unblocks main function
	done := make(chan bool)
	//user input
	for _, url := range os.Args[1:] {
		go fetch(url, ch, errCh) //start a go routine
	}
	//values of ch
	v := <-ch
	//values of errCh
	errV := <-errCh

	fmt.Printf(".%2fs elasped\n", time.Since(start).Seconds()) //runtime
}

//fetch url logic
func fetch(url string, ch chan<- string, errCh chan<- string) {
	//timer
	start := time.Now()
	//fetch url
	resp, err := http.Get(url)
	if err != nil {
		errCh <- fmt.Sprint(err) //send to channel ch
		return
	}
	//read url/close read
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		errCh <- fmt.Sprintf("while reading %s: %v", url, err) //send to channel ch
		return
	}
	secs := time.Since(start).Seconds()                  //individual url runtime
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url) //overall structure of bytes/runtime/url string
}

//write channel to file
func writeFile(v string, errCh chan<- string, done chan bool) {
	f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		errCh <- fmt.Sprintf("while opening %s: %v", f, err)
		return
	}
	defer f.Close() //close file;don't leak data
	//write to file & pass true when done
	for d := range v {
		_, err := fmt.Fprintln(f, d)
		v <- fmt.Sprintf(err)
		if err != nil {
			errCh <- fmt.Sprintf("while reading over %s: %v", v, err)
		}
	}

}

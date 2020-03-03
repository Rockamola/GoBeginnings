//fetch urls concurrently. main function runs in a goroutine and
//the go statement creates additional goroutines
package main

import (
	"bufio"
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
	//send-channel
	sendCh := make(chan string)
	//user input
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //start a go routine
	}

	//channel direction
	receiveCh(ch, sendCh)
	//
	//prints output
	//wont need print out
	for range os.Args[1:] {
		fmt.Println(<-ch) //receive from channel ch

	}

	fmt.Printf(".%2fs elasped\n", time.Since(start).Seconds()) //runtime
}

//fetch url logic
func fetch(url string, ch chan<- string) {
	//timer
	start := time.Now()
	//fetch url
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //send to channel ch
		return
	}
	//read url/close read
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err) //send to channel ch
		return
	}
	secs := time.Since(start).Seconds()                  //individual url runtime
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url) //overall structure of bytes/runtime/url string
}

//write channel to file
func writeFile(sendCh chan<- string, path string, ch chan<- string) {
	f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		ch <- fmt.Sprintf("while opening %s: %v", f, err)
		return
	}
	defer f.Close() //close file;don't leak data
	//writer w/ buffer
	w := bufio.NewWriter(f)
	//read over all lines
	for line := range  {

	}

}
func receiveCh(ch <-chan string, chSend chan<- string) {
	data := <-ch
	chSend <- data
}

//func writeFile(string <-chan ch) {

//open file
//f, err := os.Open("/home/justine/Desktop/goprojs/test.txt")
//if err != nil {
//	ch <- fmt.Sprintf("opening: %v\n", f, err)
//	f.Close()
//	return
//}
//for d := range ch {
//	_, err = fmt.Fprintln(f, d)
//}
//}

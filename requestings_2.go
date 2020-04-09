//test server or site request limit with ddos, sends requests until site is downed
//returns amounts of requested need to break server or site
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

//RequestParams values
type RequestParams struct {
	url           string
	requestsLimit int
	requestStatus *chan int

	//requests statistics
	sucessfulRequest int64
	amountRequests   int64
}

//ddos initializer for new attack/limit tester
func limitTester(URL string, requests int) (*RequestParams, error) {
	if requests < 1 {
		return nil, fmt.Errorf("Requests cannot be less than 1")
	}
	u, err := url.Parse(URL)
	if err != nil || len(u.Host) == 0 {
		return nil, fmt.Errorf("Undefined host or error =%v", err)
	}
	status := make(chan int)
	return &RequestParams{
		url:           URL,
		requestsLimit: requests,
		requestStatus: &status,
	}, nil
}

//run ddos/limit tester
func (r *RequestParams) runLimits() {
	for i := 0; i < r.requestsLimit; i++ {
		go func() {
			for {
				resp, err := http.Get(r.url)
				stat := resp.StatusCode
				if err == nil {
					_, _ = io.Copy(ioutil.Discard, resp.Body)
					_ = resp.Body.Close()
					(*r.requestStatus) <- stat
					return
				}
			}
		}()
	}
}

func main() {
	requests := 2
	r, err := limitTester("http://gopl.io", requests)
	if err != nil {
		panic(err)
	}
	r.runLimits()
	fmt.Println(r.requestStatus)
}

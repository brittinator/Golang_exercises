/*worker sends a request to the server.
If the response from the server takes too long, the
request cancels with the context.Done() functionality
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/context"
)

// global variable available throughout package
var wg sync.WaitGroup

func work(ctx context.Context) error {
	// defer always closes the wait group so it will no longer
	// block when the wait group is emptied
	defer wg.Done()

	tr := &http.Transport{}
	// HTTP client, safe to use concurrently with goroutines
	client := &http.Client{
		Transport: tr,
	}

	// struct to pack/unpack data in the channel, size of 1
	c := make(chan struct {
		r   *http.Response
		err error
	}, 1)

	// make the request to the server
	req, _ := http.NewRequest("GET", "http://localhost:1234", nil)

	go func() {
		resp, err := client.Do(req)
		fmt.Println("Doing http request is a hard(ish) job")

		pack := struct {
			r   *http.Response
			err error
		}{resp, err}

		c <- pack
	}() // must add () to execute goroutine

	// select is used when dealing with channels
	// so you can select each time if done is signaled
	select {
	// go about your business as usual
	case ok := <-c:
		if err := ok.err; err != nil {
			fmt.Println("Error ", err)
		}
		resp := ok.r
		fmt.Printf("Work work work work work... \n")

		defer resp.Body.Close()
		out, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Server Response: \n %s\n", out)

	case <-ctx.Done():
		// received the signal to cancel in this channel
		tr.CancelRequest(req)
		<-c // Wait for client.Do
		fmt.Printf("Canceled the context\n")
		return ctx.Err()
	}
	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	fmt.Println("Howdy, I'm about to do some work...")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// passing in context as argument, so if we reach the 4 second timeout
		// ctx.Done() will be kicked off and that select statement will be carried out
		go work(ctx)
	}
	wg.Wait()

	fmt.Printf("Finished. Time to go home.")
}

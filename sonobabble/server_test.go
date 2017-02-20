package sonobabble

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

/*
	TestServe tests the Serve function to see if it properly responds with a status code of 200 to a request after
	10 attempts, with an attempt every second.
*/
func TestServe(tester *testing.T) {
	// Start the server to test, with a verbose option, on a new goroutine.
	go Serve(true)

	const (
		attemptLimit uint8         = 10
		timeInterval time.Duration = time.Second
	)

	var attempt uint8
	for attempt = 0; attempt <= attemptLimit; attempt++ {
		var (
			response *http.Response
			getError error
		)
		response, getError = http.Get("http://localhost:8080")

		if getError != nil {
			tester.Fatal(getError)
		}

		/*
			Check to see if the status code is okay. If it is, exit the loop. If it is not, and it is the
			last attempt, fail the test. If neither happen, wait for the time interval before the next
			iteration.
		*/
		if response.StatusCode == http.StatusOK {
			break
		} else if response.StatusCode != http.StatusOK && attempt == attemptLimit {
			tester.Fatal(fmt.Errorf("sonobabble.testServe: recieved status code not equal to  200 after "+
				"the maximum %d attempts.", attemptLimit))
		} else {
			time.Sleep(timeInterval)
		}
	}
}

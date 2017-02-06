// NOTE: This code looks best when viewed with a tab-width of 8.

package sonobabble

import (
	"errors"
	"testing"
	"net/http"
)

/*
	TestServe tests the showHomepage function to see if it properly
	responds with a status code of 200 to a request.
*/
func TestServe(tester *testing.T) {
	// Start the server to test, with a verbose option.
	Serve(true)

	/*
		Make a pointer to the response of a get statement to the
		server, and an error if there is one.
	*/
	response, getError := http.Get("http://localhost:8080")

	// Check for any errors.
	if getError != nil {
		tester.Fatal(getError)
	}

	// Check to see if the status code is okay.
	if response.StatusCode != http.StatusOK {
		tester.Fatal(errors.New("TestServe: " +
			"recieved status code is not equal to 200"))
	}
}

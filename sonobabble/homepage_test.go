// NOTE: This code looks best when viewed with a tab-width of 8.

package sonobabble

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
	TestShowHomepage tests the showHomepage function to see if it properly responds with a status code of 200 to a
	request.
*/
func TestShowHomepage(tester *testing.T) {
	// Create a pointer to a new ResponseRecorder (which satisfies ResponseWriter for use in ServeHTTP later).
	var responseRecorder *httptest.ResponseRecorder
	responseRecorder = httptest.NewRecorder()

	// Create a HandlerFunc equal to showHomepage.
	var showHomepageHandler http.HandlerFunc
	showHomepageHandler = http.HandlerFunc(showHomepage)

	// Create a pointer to an empty request (this doesn’t really matter).
	var request *http.Request
	var requestError error
	request, requestError = http.NewRequest("GET", "", nil)

	// Check for any errors.
	if requestError != nil {
		tester.Fatal(requestError)
	}

	// Make the request.
	showHomepageHandler.ServeHTTP(responseRecorder, request)

	// Check to see if the status code is okay.
	if responseRecorder.Code != http.StatusOK {
		tester.Fatal(errors.New("sonobabble.TestShowHomepage: recieved status code is not equal to 200"))
	}

}

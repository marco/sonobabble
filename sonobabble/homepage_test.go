// NOTE: This code looks best when viewed with a tab-width of 8.

package sonobabble

import (
	"errors"
	"testing"
	"net/http"
	"net/http/httptest"
)

/*
	TestShowHomepage tests the showHomepage function to see if it properly
	responds with a status code of 200 to a request.
*/
func TestShowHomepage(tester *testing.T)  {
	/*
		Create a pointer to a new ResponseRecorder (which satisfies
		ResponseWriter for use in ServeHTTP later).
	*/
	responseRecorder := httptest.NewRecorder()

	// Create a HandlerFunc equal to showHomepage.
	showHomepageHandler := http.HandlerFunc(showHomepage)

	// Create a pointer to a new request to where Sonobabble is hosted.
	request, requestError := http.NewRequest("GET",
		"http://localhost:8080", nil)

	// Check for any errors.
	if requestError != nil {
		tester.Fatal(requestError)
	}

	// Make the request.
	showHomepageHandler.ServeHTTP(responseRecorder, request)

	// Check to see if the status code is okay.
	if responseRecorder.Code != http.StatusOK {
		tester.Fatal(errors.New("TestShowHomepage: " +
			"recieved status code is not equal to 200"))
	}

}

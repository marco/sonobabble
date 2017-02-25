package sonobabble

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
	TestShowHomepage tests the showHomepage function to see if it properly responds with a status code of 200 to a
	request.
*/
func TestShowHomepage(tester *testing.T) {
	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	var showHomepageHandler http.HandlerFunc = http.HandlerFunc(showHomepage)

	// Create an empty request (this doesn’t really matter).
	var (
		request      *http.Request
		requestError error
	)
	request, requestError = http.NewRequest("GET", "", nil)

	if requestError != nil {
		tester.Fatalf("http.NewRequest: %s", requestError)
	}

	// Make the request.
	showHomepageHandler.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		tester.Fatal("recieved status code is not equal to 200")
	}

}

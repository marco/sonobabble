package main

import (
	"fmt"

	"github.com/skunkmb/sonobabble/sonobabble"
)

// main will get the ball rolling with Sonobabble by calling the Serve function of the sonobabble package.
func main() {
	var serveError = sonobabble.Serve(true)

	/*
		Although this if statement will never be reached as long as there are no errors returned halfway
		through serving, keep it just in case.
	*/
	if serveError != nil {
		panic(fmt.Errorf("sonobabble.Serve: %s", serveError))
	}
}

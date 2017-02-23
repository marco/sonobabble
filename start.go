package main

import "github.com/skunkmb/sonobabble/sonobabble"

// main will get the ball rolling with Sonobabble by calling the Serve function of the sonobabble package.
func main() {
	// Start the server, with the verbose option.
	sonobabble.Serve(true)
}

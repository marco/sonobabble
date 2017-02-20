/*
	NOTE: This file is intended for use inside of $GOPATH/src/sonobabble (IE not the default go get location when
	downloaded from GitHub), because it imports from sonobabble/sonobabble. If you are running Sonobabble inside
	of $GOPATH/src/github.com/skunkmb/sonobabble, please use githubstart.go instead.
*/

package main

import "sonobabble/sonobabble"

// main will get the ball rolling with Sonobabble by calling the Serve function of the sonobabble package.
func main() {
	// Start the server, with the verbose option.
	sonobabble.Serve(true)
}

/*
	NOTE: This file is intended for use inside of $GOPATH/src/github.com/skunkmb/sonobabble (IE downloaded from
	GitHub), because it imports from github.com/skunkmb/sonobabble/sonobabble. If you are running Sonobabble inside
	of simply $GOPATH/src/sonobabble, please use start.go instead.
*/

package main

import "github.com/skunkmb/sonobabble/sonobabble"

// main will get the ball rolling with Sonobabble by calling the Serve function of the sonobabble package.
func main() {
	// Start the server, with the verbose option.
	sonobabble.Serve(true)
}

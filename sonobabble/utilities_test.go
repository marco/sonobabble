package sonobabble

import (
	"math/rand"
	"net/http"
	"os"
	"os/user"
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
			tester.Fatalf("http.Get http://localhost:8080: %s", getError)
		}

		/*
			Check to see if the status code is okay. If it is, exit the loop. If it is not, and it is the
			last attempt, fail the test. If neither happen, wait for the time interval before the next
			iteration.
		*/
		if response.StatusCode == http.StatusOK {
			break
		} else if response.StatusCode != http.StatusOK && attempt == attemptLimit {
			tester.Fatalf("recieved status code not equal to 200 after the maximum %d attempts",
				attemptLimit)
		} else {
			time.Sleep(timeInterval)
		}
	}
}

/*
	TestFindAbsolutePath tests the findAbsolutePath function to see if it properly returns the absolute path to a
	randomized directory inside of a temporarily-set $GOPATH.
*/
func TestFindAbsolutePath(tester *testing.T) {
	// Get the initial $GOPATH value, even if it is empty, for use later.
	var originalGoPath string = os.Getenv("GOPATH")

	var (
		testGoPath string
		pathError  error
	)
	testGoPath, pathError = generateRandomGoPath()

	if pathError != nil {
		tester.Fatalf("generateRandomGoPath: %s", pathError)
	}

	os.Setenv("GOPATH", testGoPath)

	// Create a random folder inside of a random folder (inside of $GOPATH/src/) to find the absolute path of.
	var testFolderRelativePath string = string(rand.Int()) + "/" + string(rand.Int())
	var testFolderAbsolutePath string = testGoPath + "/src/" + testFolderRelativePath
	var makeError error = os.MkdirAll(testFolderAbsolutePath, 0700)

	if makeError != nil {
		tester.Errorf("os.MkdirAll %s: %s", testFolderAbsolutePath, makeError)

		/*
			Reset $GOPATH to its original value, but don’t delete the test $GOPATH folder, because there
			was an error when creating it.
		*/
		os.Setenv("GOPATH", originalGoPath)

		tester.FailNow()
	}

	var (
		foundAbsolutePath string
		absoluteError     error
	)
	foundAbsolutePath, absoluteError = findAbsolutePath(testFolderRelativePath)

	if absoluteError != nil {
		tester.Errorf("findAbsolutePath %s: %s", testFolderRelativePath, absoluteError)

		// Reset $GOPATH to its original value, and delete the test $GOPATH folder.
		os.Setenv("GOPATH", originalGoPath)
		var removeError error = os.RemoveAll(testGoPath)

		if removeError != nil {
			// If there is an error when removing, issue an error to be shown after the first error.
			tester.Errorf("os.RemoveAll %s: %s", testFolderAbsolutePath, removeError)
		}

		tester.FailNow()
	}

	if testFolderAbsolutePath != foundAbsolutePath {
		tester.Errorf("expected %s from findAbsolutePath %s but recieved %s", testFolderAbsolutePath,
			testFolderRelativePath, foundAbsolutePath)

		// Reset $GOPATH to its original value, and delete the test $GOPATH folder.
		os.Setenv("GOPATH", originalGoPath)
		var removeError error = os.RemoveAll(testGoPath)

		if removeError != nil {
			// If there is an error when removing, issue an error to be shown after the first error.
			tester.Errorf("os.RemoveAll %s: %s", testFolderRelativePath, absoluteError)
		}

		tester.FailNow()
	}

	// Reset $GOPATH to its original value, and delete the test $GOPATH folder.
	os.Setenv("GOPATH", originalGoPath)
	var removeError error = os.RemoveAll(testGoPath)

	if removeError != nil {
		// If there is an error when removing, issue an error and fail the test.
		tester.Fatalf("os.RemoveAll %s: %s", testFolderRelativePath, absoluteError)
	}
}

/*
	generateRandomGoPath returns a random, unused path for use as a $GOPATH inside of the curerent user’s root
	directory, or an error if there is one.
*/
func generateRandomGoPath() (string, error) {
	var (
		currentUser *user.User
		userError   error
	)
	currentUser, userError = user.Current()

	if userError != nil {
		return "", userError
	}

	var newGoPath string = currentUser.HomeDir + "/" + string(rand.Int())

	var existsError error
	_, existsError = os.Stat(newGoPath)

	if existsError != nil {
		// Check to see if the error means that the directory does not exist. If it does not exist, return it.
		if os.IsNotExist(existsError) {
			return newGoPath, nil
		}

		// If there is an error, but not one about the directory not existing, return the error.
		return "", existsError
	}

	// If there is no error, that means that the directory is already taken. Therefore, try again.
	return generateRandomGoPath()
}

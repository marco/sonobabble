package sonobabble

import (
	"errors"
	"go/build"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Serve starts a Sonobabble server, and if verbose is true, outputs information on its activity.
func Serve(verbose bool) {
	// Output if needed.
	if verbose {
		log.Println("Initializing Gorilla Mux.")
	}

	// Initialize a pointer to a Gorilla Mux router.
	var router *mux.Router = mux.NewRouter()

	// Output if needed.
	if verbose {
		log.Println("Registering / pattern.")
	}

	// Use showHomepage to handle a / location.
	router.HandleFunc("/", showHomepage)

	// Get the absolute path string of the sonobabble/sonobabble package.
	var (
		absoluteSonobabblePackagePath string
		absoluteError                 error
	)
	absoluteSonobabblePackagePath, absoluteError = findAbsoluePath("sonobabble/sonobabble")

	// Check for any errors.
	if absoluteError != nil {
		panic(absoluteError)
	}

	// Create absolute path strings for three useful folders.
	var templatesAbsolutePath string = absoluteSonobabblePackagePath + "templates/"
	var templatesInstalledAbsolutePath string = absoluteSonobabblePackagePath + "templates/installed/"
	var templatesResourcesAbsolutePath string = absoluteSonobabblePackagePath + "templates/resources/"

	/*
		Create Handlers for the neccessary folders by using FileServer to return a Handler based on their
		absolute paths.

		NOTE: Dir is a string type that contains the method Open, which is required by FileServer.

		NOTE: A Handler for the templates folder is needed again here because some CSS files are stored inside
		of the templates folder, and are not loaded with HandleFunc.
	*/
	var templatesHandler http.Handler = http.FileServer(http.Dir(templatesAbsolutePath))
	var templatesInstalledHandler http.Handler = http.FileServer(http.Dir(templatesInstalledAbsolutePath))
	var templatesResourcesHandler http.Handler = http.FileServer(http.Dir(templatesResourcesAbsolutePath))

	// Output if needed.
	if verbose {
		log.Println("Registering path prefixes.")
	}

	/*
		Apply the Handlers to Routes created by PathPrefix.
	*/
	router.PathPrefix("/").Handler(templatesHandler)
	router.PathPrefix("/installed/").Handler(templatesInstalledHandler)
	router.PathPrefix("/resources/").Handler(templatesResourcesHandler)

	// Output if needed.
	if verbose {
		log.Println("Starting the server.")
	}

	// Finally, start the server.
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)

	// Output if needed.
	if verbose {
		log.Println("Finished starting the server.")
	}
}

/*
	findAbsoluePath takes in a relativePath string that is used as a package location in relation to GOPATH. It
	returns an absolute path string and an error if there is one.
*/
func findAbsoluePath(relativePath string) (string, error) {
	/*
		Check to see if a slice of strings that contains source directories for Go packages contains less than
		2 items. The first item should be GOROOT and the second item (which is needed) is GOPATH.
	*/
	if len(build.Default.SrcDirs()) < 2 {
		return "", errors.New("sonobabble.findAbsoluePath: the length of the go/build package’s default " +
			"Context’s SrcDirs slice is less than 2, therefore GOPATH is not included")
	}

	// Optain the GOPATH location
	var goPathSourceDirectory string = build.Default.SrcDirs()[1]

	/*
		Use these variables (and relativePath) to find a pointer to a Package variable of the specified
		RelativePath.
	*/
	var (
		foundPackage      *build.Package
		foundPackageError error
	)
	foundPackage, foundPackageError = build.Default.Import(relativePath, goPathSourceDirectory, build.FindOnly)

	// Check for any errors.
	if foundPackageError != nil {
		return "", foundPackageError
	}

	// Add / to the end of the Dir string of package, so that it can be properly used as a directory.
	var finalDirectory string = foundPackage.Dir + "/"

	// Return the directory.
	return finalDirectory, nil
}

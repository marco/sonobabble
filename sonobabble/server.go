// NOTE: This code looks best when viewed with a tab-width of 8.

package sonobabble

import (
	"errors"
	"github.com/gorilla/mux"
	"go/build"
	"log"
	"net/http"
)

/*
	Serve starts a Sonobabble server, and if verbose is true, outputs
	information on its activity.
*/
func Serve(verbose bool) {
	// Output if needed.
	if verbose {
		log.Println("Initializing Gorilla Mux.")
	}

	// Initialize a pointer to a Gorilla Mux router.
	router := mux.NewRouter()

	// Output if needed.
	if verbose {
		log.Println("Registering / pattern.")
	}

	// Use showHomepage to handle a / location.
	router.HandleFunc("/", showHomepage)

	/*
		Get the absolute path string of the sonobabble/sonobabble
		package.
	*/
	var absoluteSonobabblePackagePath, absoluteError =
		findAbsoluePath("sonobabble/sonobabble")

	// Check for any errors.
	if absoluteError != nil {
		panic(absoluteError)
	}

	// Create absolute path strings for three useful folders.
	var templatesAbsolutePath =
		absoluteSonobabblePackagePath + "templates/"
	var templatesInstalledAbsolutePath =
		absoluteSonobabblePackagePath + "templates/installed/"
	var templatesResourcesAbsolutePath =
		absoluteSonobabblePackagePath + "templates/resources/"

	/*
		Create Handlers for the neccessary folders by using
		FileServer to return a Handler based on their absolute
		paths. Dir is a string type that contains the method Open,
		which is required by FileServer.

		NOTE: A Handler for the templates folder is needed again here
		because some CSS files are stored inside of the templates
		folder, and are not loaded with HandleFunc.
	*/
	templatesHandler :=
		http.FileServer(http.Dir(templatesAbsolutePath))
	templatesInstalledHandler :=
		http.FileServer(http.Dir(templatesInstalledAbsolutePath))
	templatesResourcesHandler :=
		http.FileServer(http.Dir(templatesResourcesAbsolutePath))

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
	findAbsoluePath takes in a relativePath string
	that is used as a package location in relation to GOPATH. It returns
	an absolute path string and an error if there is one.
*/
func findAbsoluePath(relativePath string) (string, error) {
	// Create a pointer to the default Context.
	defaultContext := &build.Default

	/*
		Check to see if a slice of strings that contains source
		directories for Go packages contains less than 2 items. The
		first item should be GOROOT and the second item (which is
		needed) is GOPATH
	*/
	if len(defaultContext.SrcDirs()) < 2 {
		return "", errors.New("sonobabble.findAbsoluePath: the " +
			"length of the go/build package’s default Context’s " +
			"SrcDirs slice is less than 2, therefore GOPATH is " +
			"not included")
	}

	// Optain the GOPATH location
	goPathSourceDirectory := defaultContext.SrcDirs()[1]

	// Create a variable equal to the FindOnly build.ImportMode.
	findOnlyMode := build.FindOnly

	/*
		Use these variables (and relativePath) to find a pointer
		to a Package variable of the specified RelativePath.
	*/
	foundPackage, foundPackageError :=
		defaultContext.Import(relativePath,
			goPathSourceDirectory, findOnlyMode)

	// Check for any errors.
	if foundPackageError != nil {
		return "", foundPackageError
	}

	/*
		Add / to the end of the Dir string of package, so that it can
		be properly used as a directory.
	*/
	finalDirectory := foundPackage.Dir + "/"

	// Return the directory.
	return finalDirectory, nil
}

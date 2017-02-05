// NOTE: This code looks best when viewed with a tab-width of 8.

package sonobabble

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"go/build"
)

// Serve starts a Sonobabble server.
func Serve() {
	// Initialize a pointer to a Gorilla Mux router.
	router := mux.NewRouter()

	// Use showHomepage to handle a / location.
	router.HandleFunc("/", showHomepage)

	/*
		Get the absolute path string of the sonobabble/sonobabble
		package.
	*/
	var absoluteSonobabblePackagePath =
		findAbsoluePathOfPackageInGoPath("sonobabble/sonobabble")

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

	/*
		Apply the Handlers to Routes created by router.PathPrefix.
	*/
	router.PathPrefix("/").Handler(templatesHandler)
	router.PathPrefix("/installed/").Handler(templatesInstalledHandler)
	router.PathPrefix("/resources/").Handler(templatesResourcesHandler)

	// Finally, start the server.
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

/*
	findAbsoluePathOfPackageInGoPath takes in a packageRelativePath string
	that is used as a package location in relation to GOPATH. It returns
	an absolute path string.
*/
func findAbsoluePathOfPackageInGoPath(packageRelativePath string) string {
	// Create a pointer to the default Context.
	defaultContext := &build.Default

	/*
		Find the absolute path of GOPATH, assuming it is the second
		item in a slice of strings from the SrcDirs method (because
		GOROOT will be the first).
	*/
	goPathSrcDir := defaultContext.SrcDirs()[1]

	// Create a variable equal to the FindOnly build.ImportMode.
	findOnlyMode := build.FindOnly

	/*
		Use these variables (and packageRelativePath) to find a pointer
		to a Package variable of the specified packageRelativePath.
	*/
	foundPackage, foundPackageError :=
		defaultContext.Import(packageRelativePath,
		goPathSrcDir, findOnlyMode)

	// Check for any errors.
	if foundPackageError != nil {
		log.Fatal(foundPackageError)
	}

	/*
		Add / to the end of the Dir string of package, so that it can
		be properly used as a directory.
	*/
	finalDirectory := foundPackage.Dir + "/"

	// Return the directory.
	return finalDirectory
}

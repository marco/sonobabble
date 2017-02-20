package sonobabble

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Serve starts a Sonobabble server, and outputs information on its activity if using the verbose option.
func Serve(verbose bool) {
	if verbose {
		log.Println("Initializing Gorilla Mux.")
	}

	var router *mux.Router = mux.NewRouter()

	if verbose {
		log.Println("Registering patterns.")
	}

	// Use showHomepage to handle a / location, as if it were an index.html file.
	router.HandleFunc("/", showHomepage)

	/*
		Get the absolute path string of the templates folder, so that it can be used to handle all web requests
		to the templates folder on the site, such as for CSS and resources.
	*/
	var (
		templatesAbsolutePath string
		absoluteError         error
	)
	templatesAbsolutePath, absoluteError = findAbsolutePath("sonobabble/sonobabble/templates")

	if absoluteError != nil {
		panic(absoluteError)
	}

	/*
		Create a file-serving handler for the folder, which can be later applied to the / directory to host the
		templates folder on the site.
	*/
	var templatesHandler http.Handler = http.FileServer(http.Dir(templatesAbsolutePath))

	if verbose {
		log.Println("Registering path prefixes.")
	}

	// Apply the file-serving handler to the / directory.
	router.PathPrefix("/").Handler(templatesHandler)

	if verbose {
		log.Println("Starting the server.")
	}

	/*
		Finally, start the server by handling every request through the router and then listening and serving
		on port :8080.
	*/
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)

	if verbose {
		log.Println("Finished starting the server.")
	}
}

/*
	findAbsolutePath returns an absolute path that is based on a relative path in relation to the src directory in
	GOPATH, or an error if there is one. Slashes are added automatically for ease of use, so calling
	findAbsolutePath(/foo) is the same as calling findAbsolutePath(foo).

	In other words, findAbsolutePath(foo) will return the absolute path of GOPATH/src/foo.
*/
func findAbsolutePath(relativePath string) (string, error) {
	var goPath string = os.Getenv("GOPATH")

	// Join the GOPATH with the src directory and the relative path, then return the result.
	return goPath + "/src/" + relativePath, nil
}

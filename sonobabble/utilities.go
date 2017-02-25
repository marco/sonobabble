package sonobabble

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

/*
	Serve starts a Sonobabble server, outputs information on its activity if using the verbose option, and returns
	an error if there is one.
*/
func Serve(verbose bool) error {
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

		When a request is made to homepage.css or similar, it has to be served by this path, because it is not
		handled by the showHomepage function.
	*/
	var (
		templatesAbsolutePath string
		absoluteError         error
	)
	templatesAbsolutePath, absoluteError = findAbsolutePath("github.com/skunkmb/sonobabble/sonobabble/templates/")

	if absoluteError != nil {
		return fmt.Errorf("sonobabble.findAbsolutePath %s: %s",
			"github.com/skunkmb/sonobabble/sonobabble/templates/", absoluteError)
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

	// Although this return will never be reached, it is required in order to avoid an error, so return nil.
	return nil
}

/*
	findAbsolutePath returns an absolute path that is based on a relative path in relation to the src directory in
	$GOPATH, or an error if there is one. Slashes are added automatically for ease of use, so calling
	findAbsolutePath(/foo) is the same as calling findAbsolutePath(foo).

	In other words, findAbsolutePath(foo/bar) will return the absolute path of $GOPATH/src/foo/bar.
*/
func findAbsolutePath(relativePath string) (string, error) {
	var goPath string = os.Getenv("GOPATH")

	if goPath == "" {
		return "", errors.New("$GOPATH is an empty string")
	}

	var absolutePath = goPath + "/src/" + relativePath

	/*
		Ensure that the absolute path exists by using os.Stat, which returns information on a file, or an error
		if it doesn’t exist or a problem occurs.
	*/
	var existsError error
	_, existsError = os.Stat(absolutePath)

	if existsError != nil {
		return "", existsError
	}

	return absolutePath, nil
}

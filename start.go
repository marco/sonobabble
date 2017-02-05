// NOTE: This code looks best when viewed with a tab-width of 8.

package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// showHomepage writes the website’s homepage to responseWriter.
func showHomepage(responseWriter http.ResponseWriter, request *http.Request) {
	/*
		Get the absolute path of homepage.html, rather than the
		relative path.
	*/
	path, absoluteError := filepath.Abs("templates/homepage.html")

	// Then check for any errors.
	if absoluteError != nil {
		log.Fatal(absoluteError)
	}

	// Create a new pointer to a new template with the path.
	template, parseError := template.ParseFiles(path)

	// Then check for any errors.
	if parseError != nil {
		log.Fatal(parseError)
	}

	// Write the template to responseWriter.
	template.Execute(responseWriter, "")
}

func main() {
	// Initialize a pointer to a Gorilla Mux router.
	router := mux.NewRouter()

	// Use showHomepage to handle a “/” location.
	router.HandleFunc("/", showHomepage)

	/*
		Create Handlers for the templates, templates/installed, and
		templates/resources folders in two steps:

		1. Use http.Dir to get the directory.

		2. Use http.FileServer to return a Handler.

		NOTE: A Handler for the templates folder is needed again here
		because some CSS files are stored inside of the templates
		folder, and are not loaded with HandleFunc.
	*/
	templatesHandler := http.FileServer(
		http.Dir("templates/"))
	templatesInstalledHandler := http.FileServer(
		http.Dir("templates/installed/"))
	templatesResourcesHandler := http.FileServer(
		http.Dir("templates/resources/"))

	/*
		Apply the Handlers to Routes created by router.PathPrefix using
		the Handler method.
	*/
	router.PathPrefix("/").Handler(templatesHandler)
	router.PathPrefix("/installed/").Handler(templatesInstalledHandler)
	router.PathPrefix("/resources/").Handler(templatesResourcesHandler)

	// Finally, start the server.
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

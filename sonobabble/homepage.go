// NOTE: This code looks best when viewed with a tab-width of 8.

package sonobabble

import (
	"html/template"
	"net/http"
)

// showHomepage writes the website’s homepage to responseWriter.
func showHomepage(responseWriter http.ResponseWriter, request *http.Request) {
	// Get the absolute path string of the sonobabble/sonobabble package.
	var (
		absoluteSonobabblePackagePath string
		absoluteError                 error
	)
	absoluteSonobabblePackagePath, absoluteError = findAbsoluePath("sonobabble/sonobabble")

	// Check for any errors
	if absoluteError != nil {
		panic(absoluteError)
	}

	// Create the absolute path string for templates/homepage.html.
	var templatesHomepageAbsolutePath string = absoluteSonobabblePackagePath + "templates/homepage.html"

	// Create a pointer to a new Template with the path.
	var (
		template   *template.Template
		parseError error
	)
	template, parseError = template.ParseFiles(templatesHomepageAbsolutePath)

	// Check for any errors.
	if parseError != nil {
		panic(parseError)
	}

	// Write the template to responseWriter.
	template.Execute(responseWriter, "")
}

package sonobabble

import (
	"html/template"
	"net/http"
)

// showHomepage writes the website’s homepage to a response writer.
func showHomepage(responseWriter http.ResponseWriter, request *http.Request) {
	var (
		homepageAbsolutePath string
		absoluteError        error
	)
	homepageAbsolutePath, absoluteError = findAbsolutePath(
		"github.com/skunkmb/sonobabble/sonobabble/templates/homepage.html")

	if absoluteError != nil {
		panic(absoluteError)
	}

	// Create a pointer to a new template with the parsed path.
	var (
		template   *template.Template
		parseError error
	)
	template, parseError = template.ParseFiles(homepageAbsolutePath)

	if parseError != nil {
		panic(parseError)
	}

	// Write the template to the response writer.
	template.Execute(responseWriter, "")
}

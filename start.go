package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

/*
showHomepage writes the website’s homepage to responseWriter.

TODO(skunkmb): Actually write the homepage, not a placeholder.
*/
func showHomepage(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprint(responseWriter, "Welcome to Sonobabble!")
}

func main() {
	var router *mux.Router = mux.NewRouter()
	router.HandleFunc("/", showHomepage)
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

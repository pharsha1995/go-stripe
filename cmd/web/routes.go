package main

import "net/http"

// routes returns a handler to be used by http server
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	return mux
}

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	//initialize a new httprouter instance.
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	// Wrap the router with the panic recovery middleware.
	return app.recoverPanic(router)
}

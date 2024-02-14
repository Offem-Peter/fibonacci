package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/current", app.currentFibonacciHandler)
	router.HandlerFunc(http.MethodPut, "/next", app.nextFibonacciHandler)
	router.HandlerFunc(http.MethodPut, "/previous", app.prevFibonacciHandler)

	return app.recoverPanic(app.logReguest(app.rateLimit(router)))
}

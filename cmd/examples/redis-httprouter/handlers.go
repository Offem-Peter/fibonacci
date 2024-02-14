package main

import (
	"net/http"
)

func (app *application) currentFibonacciHandler(w http.ResponseWriter, r *http.Request) {
	current := app.fib.Current(ctx)

	data := envelope{
		"data": current,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) nextFibonacciHandler(w http.ResponseWriter, r *http.Request) {
	next := app.fib.Next(ctx)

	data := envelope{
		"data": next,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) prevFibonacciHandler(w http.ResponseWriter, r *http.Request) {
	prev := app.fib.Previous(ctx)

	data := envelope{
		"data": prev,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

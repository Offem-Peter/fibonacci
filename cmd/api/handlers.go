package main

import (
	"github.com/valyala/fasthttp"
)

func (app *application) currentFibonacciHandler(ctx *fasthttp.RequestCtx) {
	current := app.fib.Current(ctx)

	data := envelope{
		"data": current,
	}

	err := app.writeJSON(ctx, fasthttp.StatusOK, data)
	if err != nil {
		app.serverError(ctx, err)
	}

}

func (app *application) nextFibonacciHandler(ctx *fasthttp.RequestCtx) {
	next := app.fib.Next(ctx)

	data := envelope{
		"data": next,
	}

	err := app.writeJSON(ctx, fasthttp.StatusOK, data)
	if err != nil {
		app.serverError(ctx, err)
	}
}

func (app *application) prevFibonacciHandler(ctx *fasthttp.RequestCtx) {
	prev := app.fib.Previous(ctx)

	data := envelope{
		"data": prev,
	}

	err := app.writeJSON(ctx, fasthttp.StatusOK, data)
	if err != nil {
		app.serverError(ctx, err)
	}
}

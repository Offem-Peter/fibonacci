package main

import "github.com/valyala/fasthttp"

func (app *application) setupRoutes() fasthttp.RequestHandler {

	router := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/current":
			app.currentFibonacciHandler(ctx)
		case "/next":
			app.nextFibonacciHandler(ctx)
		case "/previous":
			app.prevFibonacciHandler(ctx)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}

	router = app.recoverPanic(router)
	router = app.logRequest(router)
	router = app.rateLimit(router)

	return router
}

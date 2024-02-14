package main

import (
	"fmt"
	"runtime/debug"

	"github.com/valyala/fasthttp"
	"golang.org/x/time/rate"
)

func (app *application) recoverPanic(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		defer func() {
			if err := recover(); err != nil {
				ctx.Response.Header.Set("Connection", "close")
				trace := fmt.Sprintf("%s\n%s", err, debug.Stack())
				app.errorLog.Output(2, trace)
				ctx.SetStatusCode(fasthttp.StatusInternalServerError)
				ctx.SetBodyString("Internal Server Error")
			}
		}()
		next(ctx)
	}
}

func (app *application) logRequest(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		app.infoLog.Printf("%s - %s %s %s",
			ctx.RemoteIP().String(),
			ctx.Response.Header.Protocol(),
			string(ctx.Method()),
			string(ctx.Path()))

		next(ctx)
	}
}

func (app *application) rateLimit(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	limiter := rate.NewLimiter(1000, 1000)

	return func(ctx *fasthttp.RequestCtx) {
		if !limiter.Allow() {
			ctx.Response.SetStatusCode(fasthttp.StatusTooManyRequests)
			return
		}
		next(ctx)
	}
}

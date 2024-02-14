package main

import (
	"encoding/json"
	"fmt"
	"runtime/debug"

	"github.com/valyala/fasthttp"
)

type envelope map[string]any

func (app *application) writeJSON(ctx *fasthttp.RequestCtx, status int, data envelope) error {
	js, err := json.Marshal(data)
	if err != nil {
		ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
		return err
	}

	ctx.Response.Header.SetContentType("application/json")
	ctx.Response.Header.SetStatusCode(status)
	ctx.Response.SetBody(js)

	return nil
}

func (app *application) serverError(ctx *fasthttp.RequestCtx, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	app.errorLog.Output(2, trace)

	ctx.Error(fasthttp.StatusMessage(fasthttp.StatusInternalServerError), fasthttp.StatusInternalServerError)
}

package main

import (
	"bytes"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func createTestContext(method, path string) *fasthttp.RequestCtx {
	ctx := &fasthttp.RequestCtx{}
	req := &fasthttp.Request{}
	req.Header.SetMethod(method)
	req.SetRequestURI(path)
	ctx.Init(req, nil, nil)

	return ctx
}

func TestRecoverPanic(t *testing.T) {
	app := &application{
		errorLog: log.New(&bytes.Buffer{}, "", log.LstdFlags),
	}

	handler := func(ctx *fasthttp.RequestCtx) {
		panic("something bad happened")
	}

	testCtx := createTestContext("GET", "/panic")

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("expected recoverPanic to handle panic, but got one: %v", r)
		}
	}()

	app.recoverPanic(handler)(testCtx)

	assert.Equal(t, fasthttp.StatusInternalServerError, testCtx.Response.StatusCode(), "Status code should be 500")

}

func TestLogRequest(t *testing.T) {
	var buf bytes.Buffer
	app := &application{
		infoLog: log.New(&buf, "", 0),
	}

	handler := func(ctx *fasthttp.RequestCtx) {}

	testCtx := createTestContext("GET", "/log")

	app.logRequest(handler)(testCtx)

	logOutput := buf.String()
	assert.Contains(t, logOutput, "GET /log", "Log output should contain the method and the path")

}

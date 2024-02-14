package main

import (
	"math/big"

	"github.com/valyala/fasthttp"
)

func (app *application) currentFibonacciHandler(ctx *fasthttp.RequestCtx) {
	fibLock.RLock()
	indexCopy := new(big.Int).Set(fib.index)
	fibLock.RUnlock()

	fibNumber := matrixFib(indexCopy)

	data := envelope{
		"data": fibNumber,
	}

	err := app.writeJSON(ctx, fasthttp.StatusOK, data)
	if err != nil {
		app.serverError(ctx, err)
	}

}

func (app *application) nextFibonacciHandler(ctx *fasthttp.RequestCtx) {
	fibLock.Lock()
	fib.index.Add(fib.index, big.NewInt(1))
	fibLock.Unlock()

	adjustedIndex := new(big.Int).Sub(fib.index, big.NewInt(1))
	fibNumber := matrixFib(adjustedIndex)

	data := envelope{
		"data": fibNumber,
	}

	err := app.writeJSON(ctx, fasthttp.StatusOK, data)
	if err != nil {
		app.serverError(ctx, err)
	}
}

func (app *application) prevFibonacciHandler(ctx *fasthttp.RequestCtx) {
	fibLock.Lock()
	if fib.index.Cmp(big.NewInt(1)) > 0 {
		fib.index.Sub(fib.index, big.NewInt(1))
	}
	fibLock.Unlock()

	adjustedIndex := new(big.Int).Sub(fib.index, big.NewInt(1))
	if adjustedIndex.Cmp(big.NewInt(0)) < 0 {
		adjustedIndex.SetInt64(0)
	}
	fibNumber := matrixFib(adjustedIndex)

	
	data := envelope{
		"data": fibNumber,
	}

	err := app.writeJSON(ctx, fasthttp.StatusOK, data)
	if err != nil {
		app.serverError(ctx, err)
	}
}

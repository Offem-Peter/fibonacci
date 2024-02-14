package main

import (
	"math/big"
	"net/http"
)

func (app *application) currentFibonacciHandler(w http.ResponseWriter, r *http.Request) {
	fibLock.RLock()
	indexCopy := new(big.Int).Set(fib.index)
	fibLock.RUnlock()

	fibNumber := matrixFib(indexCopy)

	data := envelope{
		"data": fibNumber,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) nextFibonacciHandler(w http.ResponseWriter, r *http.Request) {
	fibLock.Lock()
	fib.index.Add(fib.index, big.NewInt(1))
	fibLock.Unlock()

	adjustedIndex := new(big.Int).Sub(fib.index, big.NewInt(1))
	fibNumber := matrixFib(adjustedIndex)

	data := envelope{
		"data": fibNumber,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) prevFibonacciHandler(w http.ResponseWriter, r *http.Request) {
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

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

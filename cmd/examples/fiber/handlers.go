package main

import (
	"math/big"

	"github.com/gofiber/fiber/v2"
)

func (app *application) currentFibonacciHandler(c *fiber.Ctx) error {
	fibLock.RLock()
	indexCopy := new(big.Int).Set(fib.index)
	fibLock.RUnlock()

	fibNumber := matrixFib(indexCopy)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": fibNumber})
}

func (app *application) nextFibonacciHandler(c *fiber.Ctx) error {
	fibLock.Lock()
	fib.index.Add(fib.index, big.NewInt(1))
	fibLock.Unlock()

	adjustedIndex := new(big.Int).Sub(fib.index, big.NewInt(1))
	fibNumber := matrixFib(adjustedIndex)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": fibNumber})
}

func (app *application) prevFibonacciHandler(c *fiber.Ctx) error {
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": fibNumber})
}

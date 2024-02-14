package main

import (
	"github.com/gofiber/fiber/v2"
)

func (app *application) currentFibonacciHandler(c *fiber.Ctx) error {
	current := app.fib.Current(ctx)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": current})
}

func (app *application) nextFibonacciHandler(c *fiber.Ctx) error {
	next := app.fib.Next(ctx)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": next})
}

func (app *application) prevFibonacciHandler(c *fiber.Ctx) error {
	prev := app.fib.Previous(ctx)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": prev})
}

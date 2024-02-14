package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func (app *application) SetupRouter(fApp *fiber.App) {

	fApp.Use(logger.New())
	fApp.Use(recover.New())
	fApp.Use(rateLimit())

	fApp.Get("/current", app.currentFibonacciHandler)
	fApp.Put("/next", app.nextFibonacciHandler)
	fApp.Put("/previous", app.prevFibonacciHandler)

}

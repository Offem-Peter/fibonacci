package main

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/time/rate"
)

func rateLimit() fiber.Handler {
	limiter := rate.NewLimiter(1000, 2000)

	return func(c *fiber.Ctx) error {
		if !limiter.Allow() {
			return c.Status(fiber.StatusTooManyRequests).SendString("Too Many Requests")
		}

		return c.Next()
	}
}

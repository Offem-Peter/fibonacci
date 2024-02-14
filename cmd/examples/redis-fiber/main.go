package main

import (
	"flag"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type application struct {
	fib *FibonacciService
}

func main() {

	addr := flag.String("addr", ":8081", "HTTP network address")
	flag.Parse()

	fibService := NewFibonacciService("localhost:6381")
	fibService.writeSync = 5 * time.Second

	go fibService.syncWithRedis(ctx)

	app := &application{
		fib: fibService,
	}

	fApp := fiber.New()

	// Setup routes
	app.SetupRouter(fApp)

	// Start the server
	log.Fatal(fApp.Listen(*addr))

}

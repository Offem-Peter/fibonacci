package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
)

type application struct {
}

func main() {

	addr := flag.String("addr", ":8081", "HTTP network address")
	flag.Parse()

	app := &application{}

	fApp := fiber.New()

	app.SetupRouter(fApp)

	log.Fatal(fApp.Listen(*addr))

}

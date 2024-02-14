package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/valyala/fasthttp"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	fib      *FibonacciService
}

var app = &application{}

func main() {

	addr := flag.String("addr", ":8082", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	fibService := NewFibonacciService("localhost:6382")
	fibService.writeSync = 5 * time.Second

	go fibService.syncWithRedis(ctx)

	app = &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		fib:      fibService,
	}

	app.infoLog.Printf("Starting server on %s", *addr)

	// Start the server
	if err := fasthttp.ListenAndServe(*addr, app.setupRoutes()); err != nil {
		app.errorLog.Fatal(err)
	}

}

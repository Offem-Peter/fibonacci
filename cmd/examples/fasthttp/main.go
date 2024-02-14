package main

import (
	"flag"
	"log"
	"os"

	"github.com/valyala/fasthttp"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":8082", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	app.infoLog.Printf("Starting server on %s", *addr)

	if err := fasthttp.ListenAndServe(*addr, app.setupRoutes()); err != nil {
		app.errorLog.Fatal(err)
	}

}

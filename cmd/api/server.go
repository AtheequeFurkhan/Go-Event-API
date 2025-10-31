package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func serve(app *application) error {
	server := &http.Server{
		Addr:         fmt.Sprint(":%d", app.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Starting Server at %d", app.port)
	return server.ListenAndServe()
}

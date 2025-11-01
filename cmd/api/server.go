package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func (app *application) server() error {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Starting Server at %d", app.port)
	return server.ListenAndServe()
}

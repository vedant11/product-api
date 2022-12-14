package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/vedant11/product-api/handlers"
)

// only deals with the server setup and config
func main() {

	l := log.New(os.Stdout, "products-api ", log.LstdFlags)
	ph := handlers.NewPH(l)

	sm := mux.NewRouter()

	getMux := sm.Methods(http.MethodGet).Subrouter()
	getMux.HandleFunc("/", ph.GetProducts)
	getMux.Handle("/docs.html", http.FileServer(http.Dir("./static/")))
	getMux.Handle("/swagger.yaml", http.FileServer(http.Dir("./static/")))

	putMux := sm.Methods(http.MethodPut).Subrouter()
	putMux.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)

	postMux := sm.Methods(http.MethodPost).Subrouter()
	postMux.HandleFunc("/", ph.AddProduct)

	// create a new server
	s := http.Server{
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
		Addr:         "localhost:9090",
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}

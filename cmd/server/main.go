package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Opens channel and provides a safe way to exit channel upon termination
func createChannel() (chan os.Signal, func()) {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	return stopCh, func() {
		close(stopCh)
	}
}

// Begins http server functionality
func start(server *http.Server) {
	log.Println("Server Started")
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Failed to start server " + err.Error())
	} else {
		log.Println("Server exiting...")
	}
}

// Used after receiving interrupt to cleanly exit the program. Will force exit after 5 seconds
func shutdown(ctx context.Context, server *http.Server) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	} else {
		log.Println("Server exited")
	}
}

// uses go routine to start server.
// If an interrupt signal is received, the shutdown process begins
func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: http.FileServer(http.Dir("./assets")),
	}

	go start(server)

	// Opens channel, and waits to close it until main is returning.
	stopCh, closeCh := createChannel()
	defer closeCh()
	log.Println("Notified:", <-stopCh)

	// shut down server
	shutdown(context.Background(), server)

	// clean up channels from deferred func
}

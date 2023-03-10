package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	handler := http.FileServer(http.Dir("./assets"))
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	go func() { // start the server in a go routine so we can have the main routine listen for signals from the OS
		server.ListenAndServe()
	}()

	sigs := make(chan os.Signal, 1)                                                       // create a channel to listen for signals on
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGABRT, syscall.SIGHUP, syscall.SIGINT) // register the types of signals we're interested in

	endType := <-sigs                                            // wait for a signal notification
	fmt.Printf("received signal {%v}; shutting down\n", endType) // print the type of notification we got

	ctx, cancel := context.WithCancel(context.Background()) // create a context and get it's cancel function
	err := server.Shutdown(ctx)                             // graceful shutdown of the server
	time.Sleep(3 * time.Second)                             // wait for hopefully graceful shutdown
	cancel()                                                // cancel that graceful shutdown if it's not done
	time.Sleep(2 * time.Second)                             // wait a bit longer for the cancelled shutdown

	if err != nil { // see how the shutdown went
		fmt.Println(err.Error())
	} else {
		fmt.Println("successful shutdown of server")
	}
}

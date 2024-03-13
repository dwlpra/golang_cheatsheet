package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// create signal channel to receive signal
	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, os.Interrupt, syscall.SIGTERM)

	// app logic here
	// must inside goroutine to make it non-blocking signal terminate
	go func() {
		for {
			fmt.Println("app is running")
			time.Sleep(1 * time.Second)
		}

	}()

	// this goroutine will block until it receives a signal
	i := <-sigTerm
	fmt.Println("Got one of stop signals, SIGNAL NAME :", i)
}

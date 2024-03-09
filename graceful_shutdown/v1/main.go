package main

import (
	"context"
	"fmt"

	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const APP_TIMEOUT = time.Duration(5) * time.Second // change  time out for testing

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("worker is shutting down")
				time.Sleep(3 * time.Second) // change sleep time for testing
				fmt.Println("worker is shut down")
				return
			default:
				fmt.Println("worker is running")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	terminateSignals := make(chan os.Signal, 1)
	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGTERM)

	s := <-terminateSignals
	fmt.Println("Got one of stop signals, shutting down worker gracefully, SIGNAL NAME :", s)
	cancel()

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case <-done:
		fmt.Println("Graceful shutdown success, worker is shut down")
	case <-time.After(APP_TIMEOUT):
		fmt.Println("Graceful shutdown failed, force shutting down worker")
	}
}

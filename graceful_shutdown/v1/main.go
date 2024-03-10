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

// time out duration for graceful shutdown
const APP_TIMEOUT = time.Duration(5) * time.Second // change  time out for testing

func main() {
	// create a context that will be canceled when the app get signal to stop
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup // use wait group to controll goroutine
	wg.Add(1)
	go Worker1(ctx, &wg) // simulate worker 1

	wg.Add(1)
	go Worker2(ctx, &wg) // simulate worker 2

	wg.Add(1)
	go Worker3(ctx, &wg) // simulate worker 3

	terminateSignals := make(chan os.Signal, 1)                      // create a channel to receive signal
	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGTERM) // notify channel when get signal to stop

	s := <-terminateSignals                                                                   // block until get signal to stop
	fmt.Println("Got one of stop signals, shutting down worker gracefully, SIGNAL NAME :", s) // print signal name
	cancel()                                                                                  // cancel context to notify worker to stop

	done := make(chan struct{}) // create a channel to wait for all worker to stop
	go func() {                 // wait for all worker to stop
		wg.Wait()   // block until all worker stop
		close(done) // close done channel to notify that all worker is stopped
	}()
	select {
	case <-done: // block until all worker stop
		fmt.Println("Graceful shutdown success, worker is shut down")
	case <-time.After(APP_TIMEOUT): // when time out limit is reached, force shut down worker
		fmt.Println("Graceful shutdown failed, force shutting down worker")
	}
}

func Worker1(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker 1 is shutting down")
			time.Sleep(3 * time.Second) // change sleep time for testing
			fmt.Println("worker 1 is shut down")
			return
		default:
			fmt.Println("worker 1 is running")
			time.Sleep(1 * time.Second)
		}
	}
}

func Worker2(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker 2 is shutting down")
			time.Sleep(3 * time.Second) // change sleep time for testing
			fmt.Println("worker 2 is shut down")
			return
		default:
			fmt.Println("worker 2 is running")
			time.Sleep(1 * time.Second)
		}
	}
}

func Worker3(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker 3 is shutting down")
			time.Sleep(6 * time.Second) // change sleep time for testing
			fmt.Println("worker 3 is shut down")
			return
		default:
			fmt.Println("worker 3 is running")
			time.Sleep(1 * time.Second)
		}
	}
}

package main

import (
	"context"
	"fmt"
	"golang_cheatsheet/mock_gorm/internal/config"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/spf13/viper"
)

const APP_TIMEOUT = time.Duration(10) * time.Second

func main() {
	viperConfig := viper.New()
	viperConfig.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	db, err := config.NewDatabase(viperConfig)
	if err != nil {
		fmt.Println("Failed to connect to database")
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("worker is shutting down")
				time.Sleep(11 * time.Second)
				fmt.Println("worker is shut down")
				return
			default:
				fmt.Println("worker is running")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	terminateSignals := make(chan os.Signal, 1)
	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGTERM)

	s := <-terminateSignals
	fmt.Println("Got one of stop signals, shutting down worker gracefully, SIGNAL NAME :", s)
	cancel()

	if err := config.CloseDB(db); err != nil {
		fmt.Println("Failed to close database connection")
	}

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

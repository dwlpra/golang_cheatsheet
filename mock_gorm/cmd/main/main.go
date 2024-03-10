package main

import (
	"context"
	"fmt"
	"golang_cheatsheet/mock_gorm/internal/config"
	"golang_cheatsheet/mock_gorm/internal/delivery/routine"
	"golang_cheatsheet/mock_gorm/internal/repository"
	"golang_cheatsheet/mock_gorm/internal/usecase"

	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const APP_TIMEOUT = time.Duration(10) * time.Second

func main() {
	viperConfig := config.NewViper()

	_, cancel := context.WithCancel(context.Background())
	db, err := config.NewDatabase(viperConfig, &config.GormOpener{})
	if err != nil {
		fmt.Println("Failed to connect to database")
		return
	}
	fmt.Println("connected to database")
	var wg sync.WaitGroup

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	routine.NewUserRoutine(userUsecase)

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

package main

import (
	"context"
	"fmt"
	"github.com/IliyaYavorovPetrov/ghtmx/app/home/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	if err := Start(ctx); err != nil {
		return
	}
}

func Start(ctx context.Context) error {
	_, cancel := context.WithCancel(ctx)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handler.GetHomePage)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err.Error())
	}

	<-setupGracefulShutdown(cancel)
	return nil
}

func setupGracefulShutdown(cancel context.CancelFunc) (shutdownCompleteChan chan struct{}) {
	shutdownCompleteChan = make(chan struct{})
	isFirstShutdownSignal := true

	shutdownFunc := func() {
		if !isFirstShutdownSignal {
			log.Println("caught another exit signal, now hard dying")
			os.Exit(1)
		}

		isFirstShutdownSignal = false
		log.Println("starting graceful shutdown")

		cancel()

		close(shutdownCompleteChan)
	}

	go func(shutdownFunc func()) {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		for {
			log.Println("caught exit signal", "signal", <-sigint)
			go shutdownFunc()
		}
	}(shutdownFunc)

	return shutdownCompleteChan
}

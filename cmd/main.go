package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/IliyaYavorovPetrov/ghtmx/config"
	"github.com/IliyaYavorovPetrov/ghtmx/pkg"
	"github.com/IliyaYavorovPetrov/ghtmx/pkg/ip"
)

func main() {
	ctx := context.Background()
	if err := Run(ctx); err != nil {
		return
	}
}

func Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	cfg := config.LoadConfig()

	// init storages
	pkg.RunDatabaseSchemaMigration(cfg)
	conn := pkg.InitDatabaseConnectionPool(ctx, cfg)
	defer conn.Close()

	ipStorage := ip.NewStorage(ctx, conn)

	// init servers
	ipServer := ip.NewServer(ipStorage)

	// init handlers
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health-check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK"))
		if err != nil {
			return
		}
	})

	mux.HandleFunc("POST /ip", ip.CreateRequestIPHandler(ipServer))
	mux.HandleFunc("GET /", ip.GetRequestIPHandler(ipServer))
	mux.HandleFunc("GET /stats", ip.GetStatsIPHandler(ipServer))

	fmt.Printf("ghtmx %s is running on port %d 🚀\n", cfg.GHTMX.Version, cfg.GHTMX.Port)
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

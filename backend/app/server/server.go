package server

import (
	"context"
	"errors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sudoku/server/route"
	"time"
)

func Run(ctx context.Context) {
	mux := http.NewServeMux()
	route.Register(mux)

	// TODO: 引数で設定を受け取る
	address := "localhost:3000"
	log.Printf("Starting server on %s...", address)
	srv := &http.Server{
		Addr:    address,
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	done := make(chan error, 1)
	go func() {
		done <- srv.ListenAndServe()
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt

	select {
	case err := <-done:
		if !errors.Is(err, http.ErrServerClosed) {
			// Error starting or closing listener
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	case <-interrupt:
		log.Println("Gracefully shutting down server...")

		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout
			log.Fatalf("HTTP server Shutdown: %v", err)
		}

		log.Println("Server shutdown successfully")
	}
}

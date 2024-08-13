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
	"syscall"
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

	go func() {
		// Shutdown されると ErrServerClosed が返る
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	eg, ctx := errgroup.WithContext(context.Background())

	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintf(writer, "hello")
	})
	http.HandleFunc("/long", func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(time.Minute)
		_, _ = fmt.Fprintf(writer, "hello")
	})

	eg.Go(func() error {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	eg.Go(func() error {
		select {
		case <-signalChan:
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			return server.Shutdown(ctx)
		case <-ctx.Done():
			return nil
		}
	})

	if err := eg.Wait(); err != nil {
		log.Println("server shut down with error", err)
	} else {
		log.Println("server shut down")
	}
}

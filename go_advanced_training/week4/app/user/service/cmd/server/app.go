package main

import (
	"context"
	"errors"
	"github.com/llwwbb/geektime_practice/go_advanced_training/week4/app/user/service/internal/server"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	grpcServer *server.GrpcServer
}

func newApp(grpc *server.GrpcServer) *App {
	return &App{
		grpcServer: grpc,
	}
}

func (app *App) Run() error {
	eg, ctx := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		return app.grpcServer.Start()
	})
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	eg.Go(func() error {
		select {
		case <-signalChan:
			return app.Stop()

		case <-ctx.Done():
			return nil
		}
	})
	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (app *App) Stop() error {
	return app.grpcServer.Stop()
}

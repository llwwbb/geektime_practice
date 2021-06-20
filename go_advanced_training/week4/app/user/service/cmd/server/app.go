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
	ctx        context.Context
	cancel     context.CancelFunc
}

func newApp(grpc *server.GrpcServer) *App {
	ctx, cancel := context.WithCancel(context.Background())
	return &App{
		grpcServer: grpc,
		ctx:        ctx,
		cancel:     cancel,
	}
}

func (app *App) Run() error {
	eg, ctx := errgroup.WithContext(app.ctx)
	eg.Go(func() error {
		return app.grpcServer.Start()
	})
	eg.Go(func() error {
		<-ctx.Done()
		return app.grpcServer.Stop()
	})
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	eg.Go(func() error {
		select {
		case <-signalChan:
			return app.Stop()

		case <-ctx.Done():
			return ctx.Err()
		}
	})
	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (app *App) Stop() error {
	app.cancel()
	return nil
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/llwwbb/geektime_practice/go_advanced_training/week4/app/user/service/internal/biz"
	"github.com/llwwbb/geektime_practice/go_advanced_training/week4/app/user/service/internal/conf"
	"github.com/llwwbb/geektime_practice/go_advanced_training/week4/app/user/service/internal/data"
	"github.com/llwwbb/geektime_practice/go_advanced_training/week4/app/user/service/internal/server"
	"github.com/llwwbb/geektime_practice/go_advanced_training/week4/app/user/service/internal/service"
)

// Injectors from wire.go:

func initApp(confServer *conf.Server, confData *conf.Data) (*App, func(), error) {
	dataData, cleanup, err := data.NewData(confData)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData)
	userUseCase := biz.NewUserUseCase(userRepo)
	userService := service.NewUserService(userUseCase)
	grpcServer := server.NewGrpcServer(confServer, userService)
	app := newApp(grpcServer)
	return app, func() {
		cleanup()
	}, nil
}

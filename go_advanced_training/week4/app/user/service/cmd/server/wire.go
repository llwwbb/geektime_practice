// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/llwwbb/geektime_practice/go_advanced_training/week4/app/user/service/internal/biz"
	"github.com/llwwbb/geektime_practice/go_advanced_training/week4/app/user/service/internal/conf"
	"github.com/llwwbb/geektime_practice/go_advanced_training/week4/app/user/service/internal/data"
	"github.com/llwwbb/geektime_practice/go_advanced_training/week4/app/user/service/internal/server"
	"github.com/llwwbb/geektime_practice/go_advanced_training/week4/app/user/service/internal/service"
)

func initApp(*conf.Server, *conf.Data) (*App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}

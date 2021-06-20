package main

import (
	"flag"
	"path/filepath"

	"github.com/llwwbb/geektime_practice/go_advanced_training/week4/app/user/service/internal/conf"
	"github.com/llwwbb/geektime_practice/go_advanced_training/week4/pkg/config"
)

func main() {
	configPath := flag.String("c", "./config.yaml", "config file, eg: -c config.yaml")
	flag.Parse()
	var c conf.Config
	if err := config.LoadConfigFromFile(filepath.FromSlash(*configPath), &c); err != nil {
		panic(err)
	}
	app, cleanup, err := initApp(c.Server, c.Data)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	err = app.Run()
	if err != nil {
		panic(err)
	}
}

package data

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/llwwbb/geektime_practice/go_advanced_training/week4/app/user/service/internal/conf"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)
import "go.mongodb.org/mongo-driver/mongo"

var ProviderSet = wire.NewSet(NewData, NewUserRepo)

type Data struct {
	db *mongo.Database
}

func NewData(c *conf.Data) (*Data, func(), error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(c.MongoDB.Uri))
	if err != nil {
		return nil, nil, fmt.Errorf("mongo connect error: %w", err)
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, nil, fmt.Errorf("mongo ping error: %w", err)
	}
	db := client.Database(c.MongoDB.Db)

	return &Data{db: db}, func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Println(err)
		}
	}, nil
}

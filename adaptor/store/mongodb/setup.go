package store

import (
	"context"
	"log"

	"github.com/mohammadmokh/writino/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongodbStore struct {
	db *mongo.Database
}

func New(ctx context.Context, cfg config.MongoCfg) (MongodbStore, error) {

	clientOPs := options.Client().ApplyURI(cfg.Uri)
	client, err := mongo.Connect(ctx, clientOPs)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return MongodbStore{}, err
	}

	return MongodbStore{
		db: client.Database(cfg.DBName),
	}, nil
}

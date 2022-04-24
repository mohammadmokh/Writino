package store

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongodbStore struct {
	db *mongo.Database
}

func New(uri string) MongodbStore {

	clientOPs := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOPs)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return MongodbStore{
		db: client.Database("writino"),
	}
}

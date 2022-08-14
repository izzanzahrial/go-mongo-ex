package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection interface {
	Close()
	DB() *mongo.Database
}

type conn struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewConnection(cfg Config) (Connection, error) {
	clientOpt := options.Client().ApplyURI(cfg.Dsn())
	client, err := mongo.Connect(context.TODO(), clientOpt)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return &conn{client: client, database: client.Database(cfg.DbName())}, nil
}

func (c *conn) Close() {
	err := c.client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}

func (c *conn) DB() *mongo.Database {
	return c.database
}

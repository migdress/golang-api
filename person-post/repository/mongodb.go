package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoDBConnection(host string, port string) (*mongo.Client, error) {
	client, err := mongo.NewClient(
		options.Client().ApplyURI(
			fmt.Sprintf(
				"mongodb://%s:%s",
				host,
				port,
			),
		),
	)
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("pinging for connection..\n")
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	fmt.Printf("OK\n")
	return client, nil
}

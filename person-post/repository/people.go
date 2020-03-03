package repository

import (
	"context"
	"golang-api/person-post/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type peopleRepository struct {
	client         *mongo.Client
	ctx            context.Context
	collectionName string
	databaseName   string
	collection     *mongo.Collection
}

func (r *peopleRepository) Save(m model.Person) error {
	_, err := r.collection.InsertOne(r.ctx, m)
	if err != nil {
		return err
	}
	return nil
}

func NewPeopleRepository(
	client *mongo.Client,
	databaseName string,
	collectionName string,
	ctx context.Context,
) *peopleRepository {
	collection := client.Database(databaseName).Collection(collectionName)
	return &peopleRepository{
		client:         client,
		collection:     collection,
		collectionName: collectionName,
		databaseName:   databaseName,
	}
}

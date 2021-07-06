package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// SetIndexes will create index in db
func SetIndexes(collection *mongo.Collection, keys bsonx.Doc) error {
	unique := true
	index := mongo.IndexModel{
		Keys:    keys,
		Options: &options.IndexOptions{Unique: &unique},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if _, err := collection.Indexes().CreateOne(ctx, index); err != nil {
		return err
	}
	return nil
}

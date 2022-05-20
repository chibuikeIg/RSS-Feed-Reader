package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Ctx context.Context
}

func NewDatabase(ctxVal time.Duration) (*Database, func()) {

	ctx, cancel := context.WithTimeout(context.Background(), ctxVal*time.Second)

	return &Database{ctx}, cancel

}

func (DB Database) Client() *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI(GetEnv("DATABASE_URL")))

	if err != nil {

		log.Fatal(err)

	}

	err = client.Connect(DB.Ctx)

	return client
}

func (DB Database) Collection(cn string) *mongo.Collection {

	collection := DB.Client().Database(GetEnv("DATABASE_NAME")).Collection(cn)

	return collection
}

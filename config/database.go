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

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(GetEnv("DATABASE_URL")).
		SetServerAPIOptions(serverAPIOptions)

	client, err := mongo.Connect(DB.Ctx, clientOptions)

	if err != nil {

		log.Fatal(err)

	}

	return client
}

func (DB Database) Collection(cn string) *mongo.Collection {

	collection := DB.Client().Database(GetEnv("DATABASE_NAME")).Collection(cn)

	return collection
}

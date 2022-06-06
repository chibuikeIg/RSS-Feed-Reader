package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Ctx             context.Context
	CollectionField string
}

func NewDatabase(ctxVal time.Duration) (*Database, func()) {

	ctx, cancel := context.WithTimeout(context.Background(), ctxVal*time.Second)

	return &Database{ctx, ""}, cancel

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

func (DB Database) Find(cn string) Database {

	DB.CollectionField = cn

	return DB

}

func (DB Database) First(model any) {

	filter := bson.D{}
	opts := options.Find().SetLimit(1)
	cursor, err := DB.Collection(DB.CollectionField).Find(context.TODO(), filter, opts)
	if err = cursor.All(context.TODO(), model); err != nil {

		log.Fatal(err)

		return
	}
}

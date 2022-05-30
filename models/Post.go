package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	feed_id     primitive.ObjectID `bson:"feed_id,omitempty"`
	cover       string             `bson:"cover,omitempty"`
	title       string             `bson:"title,omitempty"`
	slug        string             `bson:"slug,omitempty"`
	description string             `bson:"description,omitempty"`
	author      string             `bson:"author,omitempty"`
	deleted_at  time.Time          `bson:"deleted_at,omitempty"`
	created_at  string             `bson:"created_at,omitempty"`
}

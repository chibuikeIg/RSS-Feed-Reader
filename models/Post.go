package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Feed_id     primitive.ObjectID `bson:"feed_id,omitempty"`
	Cover       string             `bson:"cover,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Slug        string             `bson:"slug,omitempty"`
	Description string             `bson:"description,omitempty"`
	Author      string             `bson:"author,omitempty"`
	Deleted_at  time.Time          `bson:"deleted_at,omitempty"`
	Created_at  string             `bson:"created_at,omitempty"`
}

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Feed struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	Link       string             `bson:"link,omitempty"`
	Created_at time.Time          `bson:"created_at,omitempty"`
}

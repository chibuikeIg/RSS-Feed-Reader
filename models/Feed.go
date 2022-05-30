package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Feed struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	link       string             `bson:"link,omitempty"`
	created_at time.Time          `bson:"created_at,omitempty"`
}

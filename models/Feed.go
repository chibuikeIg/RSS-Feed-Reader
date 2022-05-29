package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Feed struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	link       string             `bson:"link,omitempty"`
	created_at string             `bson:"created_at,omitempty"`
}

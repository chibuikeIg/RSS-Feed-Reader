package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Setting struct {
	Id                primitive.ObjectID `bson:"_id,omitempty"`
	Summary_length    string             `bson:"summary_length,omitempty"`
	Polling_frequency map[string]any     `bson:"polling_frequency,omitempty"`
}

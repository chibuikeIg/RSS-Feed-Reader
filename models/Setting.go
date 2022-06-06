package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Setting struct {
	Id                primitive.ObjectID `bson:"_id,omitempty"`
	summary_length    int                `bson:"summary_length,omitempty"`
	polling_frequency map[string]any     `bson:"polling_frequency,omitempty"`
}

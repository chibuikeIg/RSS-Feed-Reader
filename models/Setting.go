package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Setting struct {
	Id                primitive.ObjectID `bson:"_id,omitempty"`
	Summary_length    string             `bson:"summary_length,omitempty"`
	Polling_frequency string             `bson:"polling_frequency,omitempty"`
	Last_poll         time.Time          `bson:"last_poll,omitempty"`
}

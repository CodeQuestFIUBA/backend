package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Score is the model that represents a score object
type Score struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Room     string             `json:"room" bson:"room" validate:"required"`
	Level    string             `json:"level" bson:"level" validate:"required"`
	SubLevel string             `json:"sub_level" bson:"sub_level" validate:"required"`
	UserID   string             `json:"user_id" bson:"user_id" validate:"required"`
	Attempts int                `json:"attempts" bson:"attempts" validate:"required"`
	Points   int                `json:"points" bson:"points" validate:"required"`
}

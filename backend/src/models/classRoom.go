package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClassRoom struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      *string            `json:"name" bson:"name" validate:"required,min=2,max=100"`
	Code      *string            `json:"code" bson:"code" validate:"required,min=6"`
	Owner     string             `json:"owner" bson:"owner"`
	CreatedAt time.Time          `json:"created_at"`
	Users     []string           `json:"users" bson:"users"`
}

type ClassRoomScores struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name           *string            `json:"name" bson:"name" validate:"required,min=2,max=100"`
	Code           *string            `json:"code" bson:"code" validate:"required,min=6"`
	Owner          string             `json:"owner" bson:"owner"`
	CreatedAt      time.Time          `json:"created_at"`
	Users          []string           `json:"users" bson:"users"`
	TotalScores    int                `json:"totalScores" bson:"totalScores"`
	TotalQuestions int                `json:"totalQuestions" bson:"totalQuestions"`
}

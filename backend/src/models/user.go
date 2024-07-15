package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is the model that governs all notes objects retrived or inserted into the DB
type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	FirstName *string            `json:"first_name" validate:"required,min=2,max=100"`
	LastName  *string            `json:"last_name" validate:"required,min=2,max=100"`
	Password  *string            `json:"Password" validate:"required,min=6"`
	Email     *string            `json:"email" validate:"email,required"`
	Username  *string            `json:"username" validate:"required"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	UserId    string             `json:"user_id"`
	ClassRoom string             `json:"class_room" bson:"class_room"`
}

type UserScore struct {
	ID              primitive.ObjectID `bson:"_id"`
	FirstName       *string            `json:"first_name" validate:"required,min=2,max=100"`
	LastName        *string            `json:"last_name" validate:"required,min=2,max=100"`
	Password        *string            `json:"Password" validate:"required,min=6"`
	Email           *string            `json:"email" validate:"email,required"`
	Username        *string            `json:"username" validate:"required"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
	UserId          string             `json:"user_id"`
	ClassRoom       string             `json:"class_room" bson:"class_room"`
	Score           int                `json:"score" bson:"score"`
	TotalLevels     int                `json:"total_levels" bson:"total_levels"`
	CompletedLevels int                `json:"completed_levels" bson:"completed_levels"`
}
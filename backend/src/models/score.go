package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Score is the model that represents a score object
type Score struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ClassRoom     string             `json:"classRoom" bson:"classRoom" validate:"required"`
	Level         string             `json:"level" bson:"level" validate:"required"`
	LevelTitle    string             `json:"level_title" bson:"level_title" validate:"required"`
	SubLevel      string             `json:"sub_level" bson:"sub_level" validate:"required"`
	SubLevelTitle string             `json:"sub_level_title" bson:"sub_level_title" validate:"required"`
	UserID        string             `json:"user_id" bson:"user_id" validate:"required"`
	Attempts      int                `json:"attempts" bson:"attempts" validate:"required"`
	Points        int                `json:"points" bson:"points" validate:"required"`
}

type ScoreResponse struct {
	Level         string `json:"level"`
	Complete      bool   `json:"complete"`
	Score         int    `json:"score"`
	Qualification int    `json:"qualification"`
}

type ScoreByClassRoom struct {
	User   string `json:"user"`
	Score  int    `json:"score"`
	MyUser bool   `json:"myUser"`
}

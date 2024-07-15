package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     *string            `json:"name" validate:"required,min=2,max=100"`
	Password *string            `json:"password" validate:"required,min=6"`
	Email    *string            `json:"email" validate:"email,required"`
	AdminId  string             `json:"admin_id"`
}

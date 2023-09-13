package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    *string            `json:"firstname" validate:"required"`
	LastName     *string            `json:"lastname"`
	Password     *string            `json:"password"`
	Email        *string            `json:"email"`
	Phone        *string            `json:"phone"`
	Token        *string            `json:"token"`
	UserType     *string            `json:"usertype"`
	RefreshToken *string            `json:"refreshtoken"`
	User_id      string             `json:"userid"`
}

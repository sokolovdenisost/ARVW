package vpr

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty" db:"_id"`
	FirstName string             `json:"firstName" binding:"required" db:"firstName"`
	LastName  string             `json:"lastName" binding:"required" db:"lastName"`
	Email     string             `json:"email" binding:"required" db:"email"`
	Password  string             `json:"password" binding:"required" db:"password"`
}

type SignInBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

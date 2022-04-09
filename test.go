package vpr

import "go.mongodb.org/mongo-driver/bson/primitive"

type Test struct {
	Id    primitive.ObjectID `json:"id" bson:"_id,omitempty" db:"_id"`
	Title string             `json:"title" bson:"title" db:"title" binding:"required"`
	Tasks []Task             `json:"tasks" bson:"tasks" db:"tasks" binding:"required"`
}

type Task struct {
	Id          int    `json:"id" bson:"id" db:"id" binding:"required"`
	Description string `json:"description" bson:"description" db:"description" binding:"required"`
	Answer      string `json:"answer" bson:"answer" db:"answer" binding:"required"`
	Balls       int    `json:"balls" bson:"balls" db:"balls" binding:"required"`
}

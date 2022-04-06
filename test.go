package vpr

import "go.mongodb.org/mongo-driver/bson/primitive"

type Test struct {
	Id    primitive.ObjectID `json:"id" bson:"_id,omitempty" db:"_id"`
	Tasks []Task             `json:"tasks" bson:"tasks" db:"tasks"`
}

type Task struct {
	Description string `json:"description" bson:"description" db:"description"`
	Answer      string `json:"answer" bson:"answer" db:"answer"`
	Balls       int    `json:"balls" bson:"balls" db:"balls"`
}

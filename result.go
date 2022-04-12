package vpr

import "go.mongodb.org/mongo-driver/bson/primitive"

type Result struct {
	User    primitive.ObjectID `json:"user" bson:"user"`
	Test    primitive.ObjectID `json:"test" bson:"test,omitempty"`
	Answers []Answers          `json:"answers" bson:"answers"`
}

type Answers struct {
	TaskId int    `json:"taskId" bson:"taskId"`
	Answer string `json:"answer" bson:"answer"`
}

type ResultResponse struct {
	User    primitive.ObjectID `json:"user" bson:"user"`
	Test    []Test             `json:"tests" bson:"tests,omitempty"`
	Answers []Answers          `json:"answers" bson:"answers"`
}

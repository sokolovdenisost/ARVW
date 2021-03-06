package repository

import (
	vpr "example"

	"go.mongodb.org/mongo-driver/mongo"
)

const ErrorNotFound = "mongo: no documents in result"

type Authorization interface {
	GetUserByEmailRepo(email string) (*vpr.User, *vpr.Error)
	CreateUserRepo(reqBody vpr.User) *vpr.Error
	GetUserByIdRepo(id string) (*vpr.User, *vpr.Error)
}

type Tests interface {
	CreateTestsRepo(test vpr.Test) *vpr.Error
	GetTestsRepo() (*[]vpr.Test, *vpr.Error)
	GetTestByIdRepo(id string, answers bool) (*vpr.Test, *vpr.Error)
}

type Results interface {
	CreateResultRepo(body vpr.Result) (*string, *vpr.Error)
	GetResultsRepo(id string) (*[]vpr.ResultResponse, *vpr.Error)
}

type Repository struct {
	Authorization
	Tests
	Results
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Authorization: NewAuth(db.Collection("users")),
		Tests:         NewTests(db.Collection("tests")),
		Results:       NewResults(db.Collection("results")),
	}
}

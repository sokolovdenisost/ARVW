package repository

import (
	vpr "example"

	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	GetUserByEmailRepo(email string) (*vpr.User, *vpr.Error)
	CreateUserRepo(reqBody vpr.User) *vpr.Error
	GetUserByIdRepo(id string) (*vpr.User, *vpr.Error)
}

type Tests interface {
	CreateTestsRepo(reqBody vpr.Test) *vpr.Error
	GetTestsRepo() (*[]vpr.Test, *vpr.Error)
}

type Repository struct {
	Authorization
	Tests
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Authorization: NewAuth(db.Collection("users")),
		Tests:         NewTests(db.Collection("tests")),
	}
}

package repository

import (
	vpr "github.com/sokolovdenisost/VPR"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	GetUserByEmailRepo(email string) (*vpr.User, *vpr.Error)
	CreateUserRepo(reqBody vpr.User) *vpr.Error
	GetUserByIdRepo(id string) (*vpr.User, *vpr.Error)
}

type Tests interface {
}

type Repository struct {
	Authorization
	Tests
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Authorization: NewAuth(db.Collection("users")),
	}
}

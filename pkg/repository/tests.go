package repository

import (
	vpr "example"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TestsRepo struct {
	db *mongo.Collection
}

func NewTests(db *mongo.Collection) *TestsRepo {
	return &TestsRepo{db: db}
}

func (r *TestsRepo) CreateTestsRepo(reqBody vpr.Test) *vpr.Error {
	_, err := r.db.InsertOne(nil, reqBody)

	if err != nil {
		return SetError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (r *TestsRepo) GetTestsRepo() (*[]vpr.Test, *vpr.Error) {
	var tests []vpr.Test
	options := options.Find().SetProjection(bson.M{
		"tasks": bson.M{
			"answer": 0,
		},
	})
	cursor, err := r.db.Find(nil, bson.M{}, options)

	if err != nil {
		return nil, SetError(http.StatusInternalServerError, err.Error())
	}

	if err := cursor.All(nil, &tests); err != nil {
		return nil, SetError(http.StatusInternalServerError, err.Error())
	}

	return &tests, nil
}

package repository

import (
	vpr "example"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ResultsRepo struct {
	db *mongo.Collection
}

func NewResults(db *mongo.Collection) *ResultsRepo {
	return &ResultsRepo{db: db}
}

func (r *ResultsRepo) CreateResultRepo(body vpr.Result) *vpr.Error {
	_, err := r.db.InsertOne(nil, body)

	if err != nil {
		return SetError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (r *ResultsRepo) GetResultsRepo(id string) (*[]vpr.Result, *vpr.Error) {
	var results []vpr.Result

	objectID, errID := primitive.ObjectIDFromHex(id)

	if errID != nil {
		return nil, SetError(http.StatusInternalServerError, errID.Error())
	}

	filter := bson.M{"user": objectID}

	cursor, err := r.db.Find(nil, filter)

	if err != nil {
		return nil, SetError(http.StatusInternalServerError, err.Error())
	}

	if err := cursor.All(nil, &results); err != nil {
		return nil, SetError(http.StatusInternalServerError, err.Error())
	}

	return &results, nil
}

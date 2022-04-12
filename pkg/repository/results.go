package repository

import (
	vpr "example"
	"fmt"
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

func (r *ResultsRepo) CreateResultRepo(body vpr.Result) (*string, *vpr.Error) {
	result, err := r.db.InsertOne(nil, body)

	if err != nil {
		return nil, SetError(http.StatusInternalServerError, err.Error())
	}

	hex := fmt.Sprint(result.InsertedID)

	return &hex, nil
}

func (r *ResultsRepo) GetResultsRepo(id string) (*[]vpr.ResultResponse, *vpr.Error) {
	var results []vpr.ResultResponse

	objectID, errID := primitive.ObjectIDFromHex(id)

	if errID != nil {
		return nil, SetError(http.StatusInternalServerError, errID.Error())
	}

	matchStage := bson.D{{"$match", bson.D{{"user", objectID}}}}
	lookupStage := bson.D{{
		"$lookup", bson.D{
			{"from", "tests"},
			{"localField", "test"},
			{"foreignField", "_id"},
			{"as", "tests"},
		},
	}}

	cursor, err := r.db.Aggregate(nil, mongo.Pipeline{matchStage, lookupStage})

	if err != nil {
		return nil, SetError(http.StatusInternalServerError, err.Error())
	}

	if err := cursor.All(nil, &results); err != nil {
		return nil, SetError(http.StatusInternalServerError, err.Error())
	}

	return &results, nil
}

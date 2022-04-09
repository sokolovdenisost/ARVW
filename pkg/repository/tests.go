package repository

import (
	vpr "example"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TestsRepo struct {
	db       *mongo.Collection
	dbResult *mongo.Collection
}

func NewTests(db *mongo.Collection, dbResult *mongo.Collection) *TestsRepo {
	return &TestsRepo{
		db:       db,
		dbResult: dbResult,
	}
}

func (r *TestsRepo) CreateTestsRepo(test vpr.Test) *vpr.Error {
	fmt.Print(test)
	_, err := r.db.InsertOne(nil, test)

	if err != nil {
		return SetError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (r *TestsRepo) GetTestsRepo() (*[]vpr.Test, *vpr.Error) {
	var tests []vpr.Test
	options := options.Find().SetProjection(bson.M{
		"tasks": bson.M{
			"answer": false,
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

func (r *TestsRepo) GetTestByIdRepo(id string, answers bool) (*vpr.Test, *vpr.Error) {
	var test vpr.Test

	objectId, errID := primitive.ObjectIDFromHex(id)
	isValid := primitive.IsValidObjectID(id)

	if errID != nil {
		return nil, SetError(http.StatusInternalServerError, errID.Error())
	}

	if !isValid {
		return nil, SetError(http.StatusNotFound, "Is not a valid ObjectID")
	}

	options := options.FindOne()

	if !answers {
		options.SetProjection(bson.M{
			"tasks": bson.M{
				"answer": answers,
			},
		})
	} else {
		options.SetProjection(bson.M{})
	}

	filter := bson.M{"_id": objectId}

	err := r.db.FindOne(nil, filter, options).Decode(&test)

	if err != nil && err.Error() == ErrorNotFound {
		return nil, SetError(http.StatusNotFound, "Not found")
	}

	if err != nil {
		return nil, SetError(http.StatusInternalServerError, err.Error())
	}

	return &test, nil
}

func (s *TestsRepo) SendAnswersRepo(id string, result vpr.Result) *vpr.Error {
	return nil
}

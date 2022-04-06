package repository

import (
	"net/http"

	vpr "example"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Auth struct {
	db *mongo.Collection
}

func NewAuth(db *mongo.Collection) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateUserRepo(body vpr.User) *vpr.Error {
	_, err := r.db.InsertOne(nil, body)

	if err != nil {
		return SetError(http.StatusInternalServerError, "Bad request")
	}

	return nil
}

func (r *Auth) GetUserByEmailRepo(email string) (*vpr.User, *vpr.Error) {
	var user vpr.User

	filter := bson.M{"email": email}

	err := r.db.FindOne(nil, filter).Decode(&user)

	if user.Id.IsZero() {
		return nil, nil
	}

	if err != nil {
		return nil, SetError(http.StatusInternalServerError, err.Error())
	}

	return &user, nil
}

func (r *Auth) GetUserByIdRepo(id string) (*vpr.User, *vpr.Error) {
	var user vpr.User

	objectId, errID := primitive.ObjectIDFromHex(id)

	if errID != nil {
		return nil, SetError(http.StatusInternalServerError, "")
	}

	filter := bson.M{"_id": objectId}

	err := r.db.FindOne(nil, filter).Decode(&user)

	if user.Id.IsZero() {
		return nil, nil
	}

	if err != nil {
		return nil, SetError(http.StatusInternalServerError, err.Error())
	}

	return &user, nil
}

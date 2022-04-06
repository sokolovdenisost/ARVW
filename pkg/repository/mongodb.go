package repository

import (
	"context"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	usersTable = "users"
)

func NewClient() (*mongo.Database, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(viper.GetString("mongo_uri")))

	if err != nil {
		return nil, err
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	return client.Database("ARVW"), nil
}

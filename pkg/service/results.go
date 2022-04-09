package service

import (
	vpr "example"
	"example/pkg/repository"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ResultsService struct {
	repo repository.Results
}

func NewResultsService(repo repository.Results) *ResultsService {
	return &ResultsService{repo: repo}
}

func (s *ResultsService) CreateResultService(body vpr.Result) *vpr.Error {
	userIDok := primitive.IsValidObjectID(body.User.String())

	if userIDok {
		return SetError(http.StatusBadRequest, "Is not a valid user")
	}

	testIDok := primitive.IsValidObjectID(body.Test.String())

	if testIDok {
		return SetError(http.StatusBadRequest, "Is not a valid test")
	}

	return s.repo.CreateResultRepo(body)
}

func (s *ResultsService) GetResultsService(id string) (*[]vpr.Result, *vpr.Error) {
	return s.repo.GetResultsRepo(id)
}

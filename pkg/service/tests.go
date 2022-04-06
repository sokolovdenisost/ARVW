package service

import (
	vpr "example"
	"example/pkg/repository"
)

type TestsService struct {
	repo repository.Tests
}

func NewTestsService(repo repository.Tests) *TestsService {
	return &TestsService{repo: repo}
}

func (s *TestsService) CreateTestService(reqBody vpr.Test) *vpr.Error {
	return s.repo.CreateTestsRepo(reqBody)
}

func (s *TestsService) GetTestsService() (*[]vpr.Test, *vpr.Error) {
	return s.repo.GetTestsRepo()
}

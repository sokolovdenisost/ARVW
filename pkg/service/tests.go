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
	for idx := range reqBody.Tasks {
		reqBody.Tasks[idx].Id = idx + 1
	}

	return s.repo.CreateTestsRepo(reqBody)
}

func (s *TestsService) GetTestsService() (*[]vpr.Test, *vpr.Error) {
	return s.repo.GetTestsRepo()
}

func (s *TestsService) GetTestByIdService(id string, answers bool) (*vpr.Test, *vpr.Error) {
	return s.repo.GetTestByIdRepo(id, answers)
}

func (s *TestsService) SendAnswersService(id string, result vpr.Result) *vpr.Error {
	return s.repo.SendAnswersRepo(id, result)
}

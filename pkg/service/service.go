package service

import (
	vpr "example"
	"example/pkg/repository"
)

type Authorization interface {
	CreateUserService(reqBody vpr.User) (*vpr.User, *vpr.Error)
	GenerateTokenService(body vpr.SignInBody) (string, *vpr.Error)
	ParseTokenService(accessToken string) (string, *vpr.Error)
	GetUserByIdService(id string) (*vpr.User, *vpr.Error)
}

type Tests interface {
	CreateTestService(reqBody vpr.Test) *vpr.Error
	GetTestsService() (*[]vpr.Test, *vpr.Error)
}

type Service struct {
	Authorization
	Tests
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Tests:         NewTestsService(repos.Tests),
	}
}

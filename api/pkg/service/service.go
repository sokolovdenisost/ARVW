package service

import (
	vpr "github.com/sokolovdenisost/VPR"
	"github.com/sokolovdenisost/VPR/pkg/repository"
)

type Authorization interface {
	CreateUserService(reqBody vpr.User) (*vpr.User, *vpr.Error)
	GenerateTokenService(body vpr.SignInBody) (string, *vpr.Error)
	ParseTokenService(accessToken string) (string, *vpr.Error)
	GetUserByIdService(id string) (*vpr.User, *vpr.Error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}

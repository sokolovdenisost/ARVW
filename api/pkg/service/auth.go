package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	vpr "github.com/sokolovdenisost/VPR"
	"github.com/sokolovdenisost/VPR/pkg/repository"
)

const (
	salt     = "123mfsamuhlcxvoiy"
	jwtKey   = "my_secret_key"
	tokenTTL = 12 * time.Hour
)

func SetError(status int, message string) *vpr.Error {
	return &vpr.Error{Status: status, Message: message}
}

type tokenClaims struct {
	Id string
	jwt.StandardClaims
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUserService(reqBody vpr.User) (*vpr.User, *vpr.Error) {
	user, err := s.repo.GetUserByEmailRepo(reqBody.Email)

	if err != nil {
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	reqBody.Password = generatePasswordHash(reqBody.Password)

	return nil, s.repo.CreateUserRepo(reqBody)
}

func (s *AuthService) GenerateTokenService(reqBody vpr.SignInBody) (string, *vpr.Error) {
	user, err := s.repo.GetUserByEmailRepo(reqBody.Email)

	if err != nil {
		return "", err
	}

	if user != nil {
		hashPassword := generatePasswordHash(reqBody.Password)

		if user.Password == hashPassword {
			return generateToken(user.Id.Hex()), nil
		}

		return "", SetError(http.StatusUnauthorized, "Data is not correct")
	}

	return "", SetError(http.StatusUnauthorized, "Data is not correct")
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func generateToken(id string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	signToken, _ := token.SignedString([]byte(jwtKey))

	return signToken
}

func (s *AuthService) GetUserByIdService(id string) (*vpr.User, *vpr.Error) {
	user, err := s.repo.GetUserByIdRepo(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) ParseTokenService(accessToken string) (string, *vpr.Error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errors.New("invalid signing method")
		}

		return []byte(jwtKey), nil
	})

	if err != nil {
		return "", SetError(http.StatusInternalServerError, err.Error())
	}

	claims, ok := token.Claims.(*tokenClaims)

	if !ok {
		return "", SetError(http.StatusInternalServerError, err.Error())
	}

	return claims.Id, nil
}

package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/modaniru/api-for-users/src/repository"
	"github.com/modaniru/api-for-users/src/utils"
)

var (
	secret []byte = []byte("secret")
)

type UserService struct{
	repository *repository.Repository
	requester utils.Requester
}

func NewUserService(repository *repository.Repository, requester utils.Requester) *UserService{
	return &UserService{
		repository: repository,
		requester: requester,
	}
}

func (u *UserService) Login(token string) (string, error){
	model, err := u.requester.GetUserInfo(token)
	if err != nil {
		return "", err
	}
	id, err := u.repository.IUserRepository.Login(model)
	if err != nil {
		return "", err
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"exp": time.Now().Add(12 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})
	return t.SignedString(secret)
}





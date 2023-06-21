package service

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/modaniru/api-for-users/src/model"
	"github.com/modaniru/api-for-users/src/repository"
	"github.com/modaniru/api-for-users/src/utils"
)

var (
	secret []byte = []byte("secret")
)

type UserService struct {
	repository *repository.Repository
	requester  utils.Requester
}

func NewUserService(repository *repository.Repository, requester utils.Requester) *UserService {
	return &UserService{
		repository: repository,
		requester:  requester,
	}
}

func (u *UserService) Login(token string) (string, error) {
	model, err := u.requester.GetUserInfo(token)
	if err != nil {
		return "", err
	}
	id, err := u.repository.IUserRepository.Login(model)
	if err != nil {
		return "", err
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(12 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})
	return t.SignedString(secret)
}

func (u *UserService) ValidateJwtToken(token string) (int, error) {
	args := strings.Split(token, " ")
	if len(args) != 2 {
		return 0, errors.New("token error")
	}
	token = args[1]
	t, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("claims error")
	}
	mapClaims := map[string]interface{}(claims)
	id := int(mapClaims["id"].(float64))
	return id, nil
}

func (u *UserService) GetById(id int) (*model.User, error){
	return u.repository.IUserRepository.GetById(id)
}

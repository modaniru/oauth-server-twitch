package service

import (
	"github.com/modaniru/api-for-users/src/repository"
	"github.com/modaniru/api-for-users/src/utils"
)

type Service struct{
	IUserService
}

type IUserService interface{
	Login(token string) (string, error)
}

func NewService(repository *repository.Repository, requester utils.Requester) *Service{
	return &Service{
		IUserService: NewUserService(repository, requester),
	}
}



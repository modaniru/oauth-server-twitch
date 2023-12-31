package repository

import (
	"github.com/modaniru/api-for-users/src/model"
	"github.com/modaniru/api-for-users/src/utils"
	"gorm.io/gorm"
)

const (
	userTable = "users"
)

type Repository struct {
	IUserRepository
}

type IUserRepository interface {
	Login(user *utils.UserInfo) (int, error)
	GetById(id int) (*model.User, error)
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		IUserRepository: NewUserRepository(db),
	}
}

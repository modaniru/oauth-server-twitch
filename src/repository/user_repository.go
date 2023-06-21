package repository

import (
	"fmt"
	"time"

	"github.com/modaniru/api-for-users/src/model"
	"github.com/modaniru/api-for-users/src/utils"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Login(user *utils.UserInfo) (int, error) {
	sql := fmt.Sprintf("select id from %s where twitch_id = ?", userTable)
	var id int
	s := u.db.Raw(sql, user.Id).Scan(&id)
	err := s.Error
	if err != nil {
		return 0, err
	}
	if id == 0 {
		id, err = u.saveUser(user)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}

func (u *UserRepository) saveUser(user *utils.UserInfo) (int, error) {
	sql := fmt.Sprintf("insert into %s (username, twitch_id, image_link, registration_date) values (?, ?, ?, ?) returning id", userTable)
	year, month, day := time.Now().Date()
	var id int
	err := u.db.Raw(sql, user.DisplayName, user.Id, user.ProfileImageUrl, fmt.Sprintf("%d-%d-%d", year, month, day)).Scan(&id).Error
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *UserRepository) GetById(id int) (*model.User, error){
	sql := fmt.Sprintf("select * from %s where id = ?", userTable)
	var user model.User
	err := u.db.Raw(sql, id).Scan(&user).Error
	if err != nil{
		return nil, err
	}
	return &user, err
}
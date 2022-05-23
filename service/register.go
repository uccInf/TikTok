package service

import (
	"TikTok/dao"
	"errors"
	"gorm.io/gorm"
)

func Register(name string, password string) (dao.User, error) {
	if user, err := dao.SelectUserByName(name); errors.Is(err, gorm.ErrRecordNotFound) {
		user = dao.CreateUser(name, password)
		return user, nil
	}
	return dao.User{}, errors.New("email has been used")
}

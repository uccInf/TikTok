package service

import (
	"TikTok/dao"
	"errors"
	"gorm.io/gorm"
)

func Login(name string, password string) (*dao.User, error) {
	user, err := dao.GetUserByName(name)
	if err == nil {
		if password == user.PassWord {
			return user, nil
		}
		return user, errors.New("error password")
	}
	return user, errors.New("unregistered")
}

func Register(name string, password string) (*dao.User, error) {
	if user, err := dao.GetUserByName(name); errors.Is(err, gorm.ErrRecordNotFound) {
		user = dao.CreateUser(name, password)
		return user, nil
	}
	return nil, errors.New("email has been used")
}

func CreateToken(name string, password string) string {
	return name + password
}

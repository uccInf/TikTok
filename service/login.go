package service

import (
	"TikTok/dao"
	"errors"
)

func Login(name string, password string) (dao.User, error) {
	user, err := dao.SelectUserByName(name)
	if err == nil {
		if password == user.PassWord {
			return user, nil
		}
		return user, errors.New("error password")
	}
	return user, errors.New("unregistered")
}

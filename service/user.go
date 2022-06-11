package service

import (
	"TikTok/constdef"
	"TikTok/dao"
	"TikTok/utils"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"time"
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

func CreateToken(userId int64, userName string) string {
	maxAge := 60 * 60 * 24
	customClaims := &utils.CustomClaims{
		UserId:   userId, //用户id
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, _ := token.SignedString([]byte(constdef.SecretKey))
	return tokenString
}

func GetUserByName(name string) (*dao.User, error) {
	return dao.GetUserByName(name)
}

func AddFavorite(user *dao.User, video *dao.Video) {
	dao.AddUserFavoriteVideos(user, video)
}

func RemoveFavorite(user *dao.User, video *dao.Video) {
	dao.RemoveUserFavoriteVideos(user, video)
}

package dao

import (
	"TikTok/constdef"
	"TikTok/logger"
)

type User struct {
	UserId         int64   `json:"id,omitempty" gorm:"primaryKey"`
	Name           string  `json:"name,omitempty" gorm:"unique"`
	FollowCount    int64   `json:"follow_count,omitempty"`
	FollowerCount  int64   `json:"follower_count,omitempty"`
	IsFollow       bool    `json:"is_follow,omitempty"`
	PassWord       string  `json:"pass_word,omitempty"`
	Follows        []User  `json:"follows,omitempty" gorm:"foreignkey:UserId"`
	Followers      []User  `json:"followers,omitempty" gorm:"foreignkey:UserId"`
	FavoriteVideos []Video `json:"favorite_videos,omitempty" gorm:"many2many:user_videos"`
}

func GetUserByName(name string) (u *User, e error) {
	user := User{}
	result := DB.Table(constdef.UserTableName).
		Preload("Follows").
		Preload("Followers").
		Preload("FavoriteVideos").
		Where("name = ?", name).
		First(&user)
	// errors.Is(result.Error, gorm.ErrRecordNotFound)
	return &user, result.Error
}

func CreateUser(name string, password string) (u *User) {
	user := &User{PassWord: password, Name: name}
	DB.Table(constdef.UserTableName).Create(user)
	return user
}

func AddUserFavoriteVideos(user *User, video *Video) {
	err := DB.Model(user).Association("FavoriteVideos").Append(video)
	if err != nil {
		logger.Panic(err)
		return
	}
	AddVideoFavoriteNum(video.VideoId, video.FavoriteCount)
}

func RemoveUserFavoriteVideos(user *User, video *Video) {
	err := DB.Model(user).Association("FavoriteVideos").Delete(video)
	if err != nil {
		logger.Panic(err)
		return
	}
	RemoveVideoFavoriteNum(video.VideoId, video.FavoriteCount)
}

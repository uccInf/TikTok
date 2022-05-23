package dao

type User struct {
	Id            int64  `json:"id,omitempty" gorm:"primaryKey;unique"`
	Name          string `json:"name,omitempty" gorm:"unique"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
	PassWord      string `json:"pass_word,omitempty"`
}

func SelectUserByName(name string) (u User, e error) {
	user := User{}
	result := DB.Where("name = ?", name).First(&user)
	// errors.Is(result.Error, gorm.ErrRecordNotFound)
	return user, result.Error
}

func CreateUser(name string, password string) (u User) {
	user := User{PassWord: password, Name: name}
	DB.Create(&user)
	return user
}

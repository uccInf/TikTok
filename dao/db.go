package dao

import (
	"TikTok/constdef"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitTable(d interface{}) {
	m := DB.Migrator()
	if m.HasTable(&d) {
		return
	}

	if err := m.CreateTable(&d); err != nil {
		panic(err)
	}
}

func Init() {
	var err error
	s := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		s, constdef.UserName, constdef.PassWord, constdef.Ip, constdef.Port, constdef.DataBaseName,
	)
	fmt.Println(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//DB.DB().SetMaxOpenConns(100)
	//DB.DB().SetMaxIdleConns(10)
	if err != nil {
		panic(err)
	}
	InitTable(&User{})
	InitTable(&Video{})
	InitTable(&Comment{})
}

func GetDB() *gorm.DB {
	return DB
}

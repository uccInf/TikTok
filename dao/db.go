package dao

import (
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
	DB, err = gorm.Open(mysql.Open(
		"root:12345678@tcp(127.0.0.1:3306)/TikTok?charset=utf8mb4&parseTime=True&loc=Local",
	),
		&gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitTable(&User{})
	InitTable(&Video{})
	InitTable(&Comment{})
}

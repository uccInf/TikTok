package dao

import (
	"TikTok/constdef"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	s := "%s:%s@tcp(localhost:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		s, constdef.UserName, constdef.PassWord, constdef.Port, constdef.DataBaseName,
	)
	fmt.Println(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	//DB.DB().SetMaxOpenConns(100)
	//DB.DB().SetMaxIdleConns(10)
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&User{}, &Video{}, &Comment{})

	if err != nil {
		panic(err)
	}

}

func GetDB() *gorm.DB {
	return DB
}

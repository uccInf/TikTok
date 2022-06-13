package dao

import (
	"TikTok/constdef"
	"TikTok/logger"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	s := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		s, constdef.DBUserName, constdef.DBPassWord, constdef.DBIp, constdef.DBPort, constdef.DataBaseName,
	)
	fmt.Println(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		logger.Panic(err)
	}
	DB.AutoMigrate(&User{}, &Video{}, &Comment{})

	if err != nil {
		logger.Panic(err)
	}

}

func GetDB() *gorm.DB {
	return DB
}

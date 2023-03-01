package common

import (
	"fmt"
	"im/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//"gorm.io/gorm"
)

var DB gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	username := "root"
	database := "im"
	password := "nishisb2333"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&model.User{})

	DB = *db
	return db
}

func GetDB() *gorm.DB {
	return &DB
}

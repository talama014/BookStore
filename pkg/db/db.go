package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB() (db *gorm.DB) {
	dbName := "bookstore"
	dbUserName := "root"
	dbPassWord := "example"
	dbTCP := "127.0.0.1:3306"

	dsn := dbUserName + ":" + dbPassWord + "@tcp(" + dbTCP + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	return db
}

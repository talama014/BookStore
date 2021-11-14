package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConection struct {
	dbName     string
	dbUserName string
	dbPassWord string
	dbTCP      string
}

var AdminConn = DBConection{
	dbName:     "bookstore",
	dbUserName: "root",
	dbPassWord: "example",
	dbTCP:      "127.0.0.1:3306",
}

func GetDB(conn DBConection) (db *gorm.DB) {
	dsn := conn.dbUserName + ":" + conn.dbPassWord + "@tcp(" + conn.dbTCP + ")/" + conn.dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	return db
}

// func GetDB(dbName, dbUserName, dbPassWord, dbTCP string) (db *gorm.DB) {
// 	dsn := dbUserName + ":" + dbPassWord + "@tcp(" + dbTCP + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return
// 	}
// 	return db
// }

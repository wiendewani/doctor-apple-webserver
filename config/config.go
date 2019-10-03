package config

import (
	"Doctor-apple-Webserver/structs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/ikankuy?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect to Database")
	}

	db.AutoMigrate(structs.User{})
	db.AutoMigrate(structs.UserDetail{})
	return db
}

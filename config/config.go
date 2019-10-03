package config

import (
	"../structs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/doctor_apple?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect to Database")
	}

	db.AutoMigrate(structs.User{})
	return db
}

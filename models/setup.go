package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go-jwt-mux"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	DB = db
}

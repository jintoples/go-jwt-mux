package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(host.docker.internal:3306)/go_jwt_mux"))
	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&User{})

	DB = db
}

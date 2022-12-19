package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB() {
	dsn := "root:123456@tcp(192.168.2.124:3306)/myiot?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalln("[DB ERROR] :", err)
	}
	err = db.AutoMigrate(&DeviceBasic{}, &ProductBasic{}, &UserBasic{})
	if err != nil {
		log.Fatalln("[DB ERROR] : ", err)
	}
	DB = db
}

package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity; type:varchar(50);json:"identity"`
	Name     string `gorm:"column:Name; type:varchar(50);json:"name"`
	Password string `gorm:"column:Password; type:varchar(50);json:"password"`
}

func (table UserBasic) TableName() string {
	return "user_basic"
}

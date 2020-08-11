package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`
	Password string `gorm:"varchar(11);not null;unique"`
	Status   int    //0正常 1删除
}

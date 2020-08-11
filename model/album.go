package model

import "github.com/jinzhu/gorm"

//图片相关信息
type Album struct {
	gorm.Model
	Filepath string
	Filename string
}

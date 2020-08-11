package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title   string
	Tags    string
	Short   string
	Content string
	Author  string
}

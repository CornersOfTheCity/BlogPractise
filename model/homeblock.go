package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TagLink struct {
	TagName string
	TagUrl  string
}

type HomeBlockParam struct {
	gorm.Model
	Title      string
	Tags       []TagLink
	Short      string
	Content    string
	Author     string
	CreateTime time.Time
	//文章查看地址
	Link string
	//文章修改地址
	UpdateLink string
	DeleteLink string

	//是否登陆
	IsLogin bool
}

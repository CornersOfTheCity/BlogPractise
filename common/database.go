package common

import (
	"fmt"
	"gin.practise/ginart/model"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

//初始化数据库，给全局DB赋值
func InitDB() {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetInt("datasource.port")
	username := viper.GetString("datasource.username")
	database := viper.GetString("datasource.database")
	password := viper.GetInt("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%d@tcp(%s:%d)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("fail to connect database,err:" + err.Error() + "---" + driverName)
	}

	//关联数据库
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Article{})
	db.AutoMigrate(&model.Album{})

	DB = db
}

//获取DB实例
func GetDB() *gorm.DB {
	return DB
}

package main

import (
	"gin.practise/ginart/common"
	"gin.practise/ginart/routers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
)

func main() {
	//初始化配置
	InitConfig()
	//初始化数据库
	common.InitDB()
	//获取DB对象
	db := common.GetDB()
	//程序完成关闭数据库连接
	defer db.Close()

	//使用默认配置
	r := gin.Default()
	//初始化路由
	r = routers.InitRouter(r)
	//获取运行端口
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}

	panic(r.Run())

}

//初始化读取配置文件操作
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("InitConfig错误：" + err.Error())
	}
}

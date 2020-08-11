package controllers

import (
	"fmt"
	"gin.practise/ginart/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func TagsGet(ctx *gin.Context) {
	islogin := GetSession(ctx)
	mapData := TagsMapData("tags")
	ctx.HTML(http.StatusOK, "tags.html", gin.H{"Tags": mapData, "IsLogin": islogin})
}

func TagsPost(ctx *gin.Context) {

}

//获取所有的标签map
func TagsMapData(param string) map[string]int {
	//查询出所有标签
	db := common.GetDB()
	var list []string
	tagsMap := make(map[string]int)
	rows, _ := db.Table("articles").Where("deleted_at is null").Select(param).Rows()
	for rows.Next() {
		arg := ""
		rows.Scan(&arg)
		list = append(list, arg)
	}
	fmt.Println("List数量为：", len(list))

	for _, tag := range list {
		tagList := strings.Split(tag, "&")
		for _, value := range tagList {
			tagsMap[value]++
		}
	}
	return tagsMap
}

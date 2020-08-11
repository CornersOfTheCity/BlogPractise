package controllers

import (
	"gin.practise/ginart/common"
	"gin.practise/ginart/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var article model.Article

func UpdateArticleGet(ctx *gin.Context) {
	islogin := GetSession(ctx)
	idStr := ctx.Query("id")
	id, _ := strconv.Atoi(idStr)
	db := common.GetDB()
	article = model.Article{}
	db.Where("ID  = ?", id).First(&article)
	ctx.HTML(http.StatusOK, "write_article.html", gin.H{"IsLogin": islogin, "Title": article.Title, "Tags": article.Tags, "Short": article.Short, "Content": article.Content, "Id": article.ID})
}

func UpdateArticlePost(ctx *gin.Context) {

	db := common.GetDB()
	//获取浏览器传输的数据，通过表单的name属性获取值
	article.Title = ctx.PostForm("title")
	article.Tags = ctx.PostForm("tags")
	article.Short = ctx.PostForm("short")
	article.Content = ctx.PostForm("content")
	db.Save(&article)

	//返回给浏览器
	//response := gin.H{}
	response := gin.H{"code": 1, "message": "更新成功"}
	ctx.JSON(http.StatusOK, response)
}

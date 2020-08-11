package controllers

import (
	"gin.practise/ginart/common"
	"gin.practise/ginart/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddArticleGet(ctx *gin.Context) {
	//是否为登陆状态
	islogin := GetSession(ctx)
	ctx.HTML(http.StatusOK, "write_article.html", gin.H{"IsLogin": islogin})
}

func AddArticlePost(ctx *gin.Context) {
	DB := common.GetDB()
	//获取表单细信息
	title := ctx.PostForm("title")
	tags := ctx.PostForm("tags")
	short := ctx.PostForm("short")
	content := ctx.PostForm("content")

	newArticle := model.Article{
		Title:   title,
		Tags:    tags,
		Short:   short,
		Content: content,
		Author:  "user",
	}
	DB.Create(&newArticle)
	response := gin.H{}

	ctx.JSON(http.StatusOK, response)

}

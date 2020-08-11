package controllers

import (
	"gin.practise/ginart/common"
	"gin.practise/ginart/model"
	"gin.practise/ginart/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//显示文章详情页
func ShowArticleGet(ctx *gin.Context) {
	islogin := GetSession(ctx)

	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	//根据ID查询文章
	db := common.GetDB()
	var article model.Article
	db.Where("ID  = ?", id).First(&article)
	ctx.HTML(http.StatusOK, "show_article.html", gin.H{"IsLogin": islogin, "Title": article.Title, "Content": utils.SwitchMarkdownToHtml(article.Content)})
	//ctx.HTML(http.StatusOK, "show_article.html", gin.H{"IsLogin": islogin, "Title": article.Title, "Content": article.Content})
}

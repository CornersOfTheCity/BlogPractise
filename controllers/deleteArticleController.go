package controllers

import (
	"gin.practise/ginart/common"
	"gin.practise/ginart/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteArticleController(ctx *gin.Context) {
	idStr := ctx.Query("id")
	id, _ := strconv.Atoi(idStr)
	db := common.GetDB()
	var article model.Article
	db.Where("ID  = ?", id).First(&article)
	db.Delete(&article)
	ctx.Redirect(http.StatusMovedPermanently, "/")

}

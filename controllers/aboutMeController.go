package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AboutMeGet(ctx *gin.Context) {
	islogin := GetSession(ctx)
	ctx.HTML(http.StatusOK, "aboultme.html", gin.H{"IsLogin": islogin, "wechat": "微信：12345678", "qq": "QQ：12345678", "tel": "Tel：12345678"})
}

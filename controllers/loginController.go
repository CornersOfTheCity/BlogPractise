package controllers

import (
	"gin.practise/ginart/common"
	"gin.practise/ginart/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func LoginGet(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{"title": "登陆页"})
}

func LoginPost(ctx *gin.Context) {
	db := common.GetDB()
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if len(username) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "massage": "姓名不能为空\n"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "massage": "密码不能少于六位\n"})
		return
	}

	var user model.User
	//db.Debug().Where("name=? AND password=?", username, hasedPassword).First(&user)
	db.Where("name=?", username).First(&user)
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "登陆失败"})
		return
	} else {

		session := sessions.Default(ctx)
		session.Set("loginUser", username)
		session.Save()
		/*
			//生成Token
			token, err := common.ReleaseToken(user.Name)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "系统异常"})
				return
			}
		*/
		ctx.JSON(http.StatusOK, gin.H{"code": 1, "message": "登陆成功"})
		//ctx.JSON(http.StatusOK, gin.H{"code": 1, "data": gin.H{"token": token}, "message": "登陆成功"})
	}
}

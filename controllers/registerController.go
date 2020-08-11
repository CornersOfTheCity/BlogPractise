package controllers

import (
	"fmt"
	"gin.practise/ginart/common"
	"gin.practise/ginart/model"
	"gin.practise/ginart/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func RegisterGet(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "注册页"})
}

func RegisterPost(ctx *gin.Context) {

	fmt.Println("开始注册验证")
	db := common.GetDB()
	//var requestUser = model.User{}
	//ctx.Bind(&requestUser)
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	//username := requestUser.Name
	//password := requestUser.Password
	//数据验证
	if len(password) < 6 {
		ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "密码不能少于六位"})
		return
	}

	var user model.User
	//姓名为空给一个十位随机字母
	if len(username) == 0 {
		for {
			username = utils.RandomString(10)
			db.Where("name=?", username).First(&user)
			if user.ID > 0 {
			} else {
				break
			}
		}
	} else {
		db.Where("name=?", username).First(&user)
		if user.ID > 0 {
		} else {
			ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "用户名重复"})
			return
		}
	}

	//加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		//response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误\n")
		ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "加密错误"})
		return
	}

	newUser := model.User{
		Name:     username,
		Password: string(hashedPassword),
		Status:   0,
	}

	db.Create(&newUser)
	//response.Success(ctx, nil, "注册成功\n")
	ctx.JSON(http.StatusOK, gin.H{"code": 1, "message": "注册成功"})

}

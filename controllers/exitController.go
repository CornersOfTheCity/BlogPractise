package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExitGet(ctx *gin.Context) {
	//清除登陆状态数据
	fmt.Println("调用ExitGet，，，")
	session := sessions.Default(ctx)
	session.Delete("loginUser")
	//session.Clear()
	session.Save()
	fmt.Println("delete session...")
	//清除session重定位
	//使用http状态码301永久重定向会出现当第一退出登陆之后，cookie可以正常清理，但是当第二次之后登录系统无法退出登录，修改为302解决
	ctx.Redirect(http.StatusMovedPermanently, "/")
}

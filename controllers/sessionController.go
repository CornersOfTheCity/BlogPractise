package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetSession(ctx *gin.Context) bool {

	session := sessions.Default(ctx)
	loginUser := session.Get("loginUser")
	fmt.Println("调用GetSession：", loginUser)
	if loginUser != nil {
		return true
	} else {
		return false
	}
}

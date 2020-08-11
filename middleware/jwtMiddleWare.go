package middleware

import (
	"gin.practise/ginart/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//验证JWT并将结果存储到上下文中
func JwtMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			ctx.Abort()
			return
		}
		authHeader = authHeader[7:]

		token, claims, err := common.ParseToken(authHeader)

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusOK, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//验证通过后 获取claim中的username
		userName := claims.UserName

		// 将当前请求的username信息保存到请求的上下文c上
		ctx.Set("username", userName)
		ctx.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息

	}
}

package routers

import (
	"gin.practise/ginart/controllers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	r.LoadHTMLGlob("views/*")
	r.Static("/static", "./static")

	//设置session中间件
	store := cookie.NewStore([]byte("loginUser"))
	r.Use(sessions.Sessions("mySession", store))

	//注册
	r.GET("/register", controllers.RegisterGet)
	r.POST("/register", controllers.RegisterPost)

	//登陆
	r.GET("/login", controllers.LoginGet)
	r.POST("/login", controllers.LoginPost)

	//进入首页
	r.GET("/", controllers.HomeGet)

	//退出登录
	r.GET("/exitGet", controllers.ExitGet)

	//路由组
	v1 := r.Group("/article")
	{
		//添加文章
		v1.GET("/add", controllers.AddArticleGet)
		v1.POST("/add", controllers.AddArticlePost)

		//修改文章
		v1.GET("/update", controllers.UpdateArticleGet)
		v1.POST("/update", controllers.UpdateArticlePost)

		//删除文章
		v1.GET("/delete", controllers.DeleteArticleController)
	}

	//显示文章详情
	V2 := r.Group("/show")
	{
		V2.GET("/:id", controllers.ShowArticleGet)
	}

	//获取标签
	r.GET("/tags", controllers.TagsGet)

	//相册获取
	r.GET("/album", controllers.AlbumGet)
	//图片上传
	r.POST("/upload", controllers.UploadPost)

	//关于我
	r.GET("/aboutme", controllers.AboutMeGet)

	return r
}

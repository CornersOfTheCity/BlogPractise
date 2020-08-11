package controllers

import (
	"bytes"
	"fmt"
	"gin.practise/ginart/common"
	"gin.practise/ginart/model"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

//存储表行数，作为一个只有自己可以更改的量，用于分页操作
var artNum = 0

//***************查询文章********************
func HomeGet(ctx *gin.Context) {
	//判断用户是否登陆
	islogin := GetSession(ctx)
	/*
		session := sessions.Default(ctx)
		session.Clear()
		session.Save()
	*/
	//tag有值为标签获取
	//page有值为翻页

	//page := 1

	//获取前端参数
	tag := ctx.Query("tag")
	pag, _ := strconv.Atoi(ctx.Query("page"))

	var artList []*model.Article
	var hasFooter bool

	if len(tag) > 0 {
		artList = FindArticleWithTag(tag)
		hasFooter = false
	} else {
		if pag <= 0 {
			pag = 1
		}
		artList, _ = FindArticleWithPage(pag)
		hasFooter = true
	}

	html := MakeHomeBlocks(artList, islogin)
	homeFooterPageCode := GetPages(pag)

	ctx.HTML(http.StatusOK, "home.html", gin.H{"IsLogin": islogin, "Content": html, "HasFooter": hasFooter, "PageCode": homeFooterPageCode})
}

//根据tag模糊查找相关文章
func FindArticleWithTag(tag string) []*model.Article {
	var articleList []*model.Article
	db := common.GetDB()
	db.Where("tags LIKE ?", tag).Find(&articleList)
	return articleList
}

//gorm分页查询，传入页数，偏移量，返回查询到的对象，页数，错误
func FindArticleWithPage(page int) ([]*model.Article, error) {
	//获取数据库操作对象
	var DB = common.GetDB()
	var articles []*model.Article
	//var count uint64

	pagSize := viper.GetInt("pages.num")

	//tdb.Limit(pagSize).Offset(pagSize * (page - 1))
	//fmt.Println("获取偏移数据")

	//使用count  不能用 Offset 或将Offset值设为 -1（-1代表取消offset限制）

	//if err := DB.Table("articles").Where("deleted_at is null").Limit(pagSize).Offset(pagSize * (page - 1)).Find(&articles).Count(&count).Error; err != nil
	if err := DB.Table("articles").Where("deleted_at is null").Limit(pagSize).Offset(pagSize * (page - 1)).Find(&articles).Error; err != nil {
		fmt.Println(err.Error())
		return articles, err
	}
	fmt.Println("完成")

	return articles, nil
}

//将文章内容显示到页面上
func MakeHomeBlocks(articles []*model.Article, isLogin bool) template.HTML {
	htmlHome := ""
	for _, art := range articles {
		//将数据库model转换为首页模板所需要的model
		homeParam := model.HomeBlockParam{}
		homeParam.ID = art.ID
		homeParam.Title = art.Title
		homeParam.Tags = CreateTagsLinks(art.Tags)
		//fmt.Println("tag-->", art.Tags)
		homeParam.Short = art.Short
		homeParam.Content = art.Content
		homeParam.Author = art.Author
		homeParam.CreateTime = art.CreatedAt
		homeParam.Link = "/show/" + strconv.Itoa(int(art.ID))
		homeParam.UpdateLink = "/article/update?id=" + strconv.Itoa(int(art.ID))
		homeParam.DeleteLink = "/article/delete?id=" + strconv.Itoa(int(art.ID))
		homeParam.IsLogin = isLogin

		//处理变量
		//ParseFile解析该文件，用于插入变量
		t, _ := template.ParseFiles("views/home_block.html")
		buffer := bytes.Buffer{}
		//就是将html文件里面的比那两替换为穿进去的数据
		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	//fmt.Println("htmlHome-->", htmlHome)
	return template.HTML(htmlHome)
}

//将tags字符串转化成首页模板所需要的数据结构
func CreateTagsLinks(tags string) []model.TagLink {
	var tagLink []model.TagLink
	tagsPamar := strings.Split(tags, "&")
	for _, tag := range tagsPamar {
		tagLink = append(tagLink, model.TagLink{tag, "/?tag=" + tag})
	}
	return tagLink
}

//****************分页**********************
//根据分页后页数获取的Pages对象
func GetPages(page int) model.Pages {
	pageCode := model.Pages{}
	//查询出总条数
	num := GetArtRowNum()
	//计算总页数
	pag := viper.GetInt("pages.num")
	allPage := (num-1)/pag + 1
	pageCode.ShowPage = fmt.Sprintf("%d/%d", page, allPage)
	//当前页数小于1则没有上一页按钮
	if page <= 1 {
		pageCode.HasPre = false
	} else {
		pageCode.HasPre = true
	}

	//当前页数大于总页数则没有下一页
	if page >= allPage {
		pageCode.HasNext = false
	} else {
		pageCode.HasNext = true
	}

	pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)
	pageCode.NextLink = "/?page=" + strconv.Itoa(page+1)
	return pageCode
}

//获取总文章数目
func GetArtRowNum() int {
	//获取数据库操作对象
	var DB = common.GetDB()
	if artNum == 0 {
		DB.Table("articles").Where("title <> ?", "daw").Count(&artNum)
	}
	return artNum
}

//设置页数,用于添加或删除文章时更新页数
func SetArtNum() {
	artNum = GetArtRowNum()
}

//添加文章
func AddArticle(article model.Article) (uint, error) {
	//获取数据库操作对象
	var DB = common.GetDB()
	DB.Create(&article)
	return article.ID, nil
}

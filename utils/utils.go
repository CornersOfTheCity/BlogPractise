package utils

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"github.com/sourcegraph/syntaxhighlight"
	"html/template"
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letters = []byte{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p', 'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'z', 'x', 'c', 'v', 'b', 'n', 'm'}
	rand.Seed(time.Now().Unix())
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

//将传入的时间戳转为时间
func SwitchTimeStampToData(timeStamp int64) string {
	t := time.Unix(timeStamp, 0)
	return t.Format("2006-01-02 15:04:05")
}

//利用MarkDown编辑文档
func SwitchMarkdownToHtml(content string) template.HTML {
	markdown := blackfriday.MarkdownCommon([]byte(content))

	//获取HTML文档
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(markdown))
	/*
	  对document进程查询，选择器和css的语法一样
	  第一个参数：i是查询到的第几个元素
	  第二个参数：selection就是查询到的元素
	*/
	doc.Find("code").Each(func(i int, selection *goquery.Selection) {
		light, _ := syntaxhighlight.AsHTML([]byte(selection.Text()))
		selection.SetHtml(string(light))
		//fmt.Println(selection.Html())
		//fmt.Println("light:", string(light))
		//fmt.Println("\n\n\n")
	})
	htmlString, _ := doc.Html()
	return template.HTML(htmlString)
}

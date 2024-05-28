package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Game struct {
	ChineseName string `json:"chineseName"` //中文名称
	EnglishName string `json:"englishName"` //英文名称
	//Price       float64 `json:"price"`       //价格
	CoverURL string `json:"coverURL"` //封面URL
}
type Item struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := gin.Default()
	// 告诉gin框架模板文件中的静态文件的位置
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	// 查找过程
	r.GET("/search", SearchWhenInput)
	r.GET("/search/:englishName", func(c *gin.Context) {})
	r.Run()
}

func SearchWhenInput(context *gin.Context) {
	query, ok := context.GetQuery("q")
	if !ok {
		context.JSON(http.StatusServiceUnavailable, gin.H{
			"code": http.StatusServiceUnavailable,
			"msg":  "服务异常",
		})
	}
	var items []Item
	results := getName(query)
	for _, item := range results {
		items = append(items, Item{
			Name: item,
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"data": items,
	})
}

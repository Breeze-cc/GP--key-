package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
)

type Game struct {
	ChineseName string  //中文名称
	EnglishName string  //英文名称
	Price       float64 //价格
	Unit        string  //货币单位
	Cover       string  //封面URL
}

func main() {
	r := gin.Default() // 初始化gin
	r.Static("/static", "./statics")
	r.SetFuncMap(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	})
	r.LoadHTMLGlob("statics/templates/**/*")
	r.GET("/game_price/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "game/index.tmpl", gin.H{ //模板渲染
			"title": "<h1>字符串</h1>",
		})
	})
	r.GET("/json", func(c *gin.Context) {
		m := Game{
			"传说之下",
			"UnderTale",
			13.4,
			"元",
			"www",
		}
		c.JSON(http.StatusOK, m)
	})
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("run server error", err)
	}
}

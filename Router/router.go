package Router

import (
	"elasticsearch/Router/Funs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RunGin(){

	router := gin.Default()

	//Service路由组
	Books := router.Group("/books")
	{
		Books.Handle("POST","/search",Funs.BoolByBooks)
	}

	Press:= router.Group("/helper")
	{
		Press.Handle("GET","/press",Funs.PressList)

		Press.Handle("GET","/load",Funs.LoadBook)
		Press.Handle("GET","/press:press",Funs.LoadBookByPress)
		Press.Handle("GET","/terms:press",Funs.TermsBookPress)
		Press.Handle("GET","/match:name",Funs.MatchBookName)
		Press.Handle("GET","/source/collapse",Funs.PressList)
		Press.Handle("GET","/Range",Funs.RangeQueryBook)
	}

	//vue前端页面
	router.StaticFS("/ui",http.Dir("./htmls"))
	router.Run(":8080")
}

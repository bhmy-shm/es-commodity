package Funs

import (
	"elasticsearch/AppInit"
	"elasticsearch/Model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"strings"
)

//1.Search查询默认十条
func LoadBook(ctx *gin.Context){
	rest,err := AppInit.ConnEs().Search().Index("books").Do(ctx)
	if err != nil{
		ctx.JSON(500,gin.H{"error":err})
	}else{
		ctx.JSON(200,gin.H{"result":ResultToBooks(rest)})
	}
}


//2.term精确查找,查找 BookPress:"湘潭大学出版社"
func LoadBookByPress(ctx *gin.Context){
	press,_ := ctx.Params.Get("press")

	termQuery := elastic.NewTermQuery("BookPress",press)
	rest,err := AppInit.ConnEs().Search().Query(termQuery).Index("books").Do(ctx)
	if err != nil{
		ctx.JSON(500,gin.H{"error":err})
	}else{
		ctx.JSON(200,gin.H{"result":ResultToBooks(rest)})
	}
}


//3.terms多精确查找
func TermsBookPress(ctx *gin.Context){
	//注意这里的从ctx拿到的是多个 terms条件，所以要用逗号分开
	press,_ := ctx.Params.Get("press")
	list := strings.Split(press,",")

	PresSlice := []interface{}{}
	for _,p :=range list {
		PresSlice = append(PresSlice,p)
	}

	//创建Terms入口
	termsQuery := elastic.NewTermsQuery("BookPress",PresSlice...)
	rest,err := AppInit.ConnEs().Search().Query(termsQuery).Index("books").Size(30).Do(ctx)
	if err != nil{
		ctx.JSON(500,gin.H{"error":err})
	}else{
		ctx.JSON(200,gin.H{"result":ResultToSource(rest)})
	}
}

//4.全文高亮检索
func MatchBookName(ctx *gin.Context){
	name,_ := ctx.Params.Get("name")
	machQuery := elastic.NewMatchQuery("BookName",name)
	field_tag := elastic.NewHighlight().Field("BookName").PreTags("<t>").PostTags("</t>")

	rest,err := AppInit.ConnEs().Search().Query(machQuery).Index("books").
		Highlight(field_tag).Do(ctx)
	if err != nil{
		ctx.JSON(500,gin.H{"error":err})
	}else{
		ctx.JSON(200,gin.H{"result":ResultToSource(rest)})
	}
}

//5.Source去重查询
func PressList(ctx *gin.Context) {
	cb := elastic.NewCollapseBuilder("BookPress")
	rest,err := AppInit.ConnEs().Search().Collapse(cb).FetchSource(false).
		Index("books").Size(20).Do(ctx)
	if err != nil{
		ctx.JSON(500,gin.H{"error":err})
	}else{
		ctx.JSON(200,gin.H{"result":ResultToFileds(rest,"BookPress")})
	}
}

//6.RangeQuery 范围查询
func RangeQueryBook(ctx *gin.Context){
	//
	search := Model.NewSearchModel()
	err := ctx.BindJSON(&search)
	fmt.Println("111",search.BookPrice1Start,search.BookPrice2End)
	if err != nil {
		ctx.JSON(500,gin.H{"error":err})
	}
	//如果起始和末尾范围都有值，就创建Range查询
	if search.BookPrice1Start > 0 || search.BookPrice2End >0 {
		Rangequery := elastic.NewRangeQuery("BookPrice1")

		//查询的起始范围比结束范围小才能执行RangeQuery。
		//如果 起始范围还有没结束范围 大，则报错
		if search.BookPrice1Start < search.BookPrice2End {
			Rangequery.Gte(search.BookPrice1Start) //起始范围
			Rangequery.Lte(search.BookPrice2End)   //结束范围

			rest,err := AppInit.ConnEs().Search().Query(Rangequery).Index("books").Do(ctx)
			if err != nil{
				ctx.JSON(500,gin.H{"error":err})
			}else{
				ctx.JSON(200,gin.H{"result":rest})
			}
		}else{
			ctx.JSON(506,gin.H{"error":"start or end is failed"})
		}
	}
}

//
package Funs

import (
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"strings"
	"elasticsearch/AppInit"
	"elasticsearch/Model"
)

//vue前端搜索页面
func BoolByBooks(ctx *gin.Context){

	//1.获取POST的提交参数
	search := Model.NewSearchModel()
	err := ctx.BindJSON(&search)
	if err != nil {
		ctx.JSON(500,gin.H{"error":err})
	}

	//2.创建 query查询切片
	Querys := make([]elastic.Query,0)

	//Match 查询书名 追加 Query切片
	if search.BookName != ""{
		match := elastic.NewMatchQuery("BookName",search.BookName)
		Querys = append(Querys,match)
	}
	//Term 查询出版社 追加 Query切片
	if search.BookPress != ""{

		press := strings.Split(search.BookPress,",")
		presslist := []interface{}{}
		for _,p := range press {
			presslist = append(presslist,p)
		}
		terms := elastic.NewTermsQuery("BookPress",presslist...)
		Querys = append(Querys,terms)
	}
	//RangeQuery范围查找金额 追加 Query切片
	if search.BookPrice1Start > 0 || search.BookPrice2End >0 {
		Rangequery := elastic.NewRangeQuery("BookPrice1")

		//查询的起始范围比结束范围小才能执行RangeQuery。
		//如果 起始范围还有没结束范围 大，则报错
		if search.BookPrice1Start < search.BookPrice2End {
			Rangequery.Gte(search.BookPrice1Start) //起始范围
			Rangequery.Lte(search.BookPrice2End)   //结束范
		}
		Querys = append(Querys,Rangequery)
	}
	//Sort 排序 追加 Query切片
	sortList := make([]elastic.Sorter,0)
	{
		//如果需要score精度排序
		if search.OrderSet.Score{
			sortList = append(sortList,elastic.NewScoreSort().Desc())
		}
		//如果是ASC 从小到大排序
		if search.OrderSet.PriceOrder == Model.OrderByPriceASC{
			sortList = append(sortList, elastic.NewFieldSort("BookPrice1").Asc())
		}
		//如果是DESC，从高到低排序
		if search.OrderSet.PriceOrder == Model.OrderByPriceDESC{
			sortList = append(sortList, elastic.NewFieldSort("BookPrice1").Desc())
		}
	}

	//4.创建一个 Bool查询，并设定Bool的Must参数，将query切片追加到Must当中
	BoolMustQuery := elastic.NewBoolQuery().Must(Querys...)

	//5.搜索Bool的结果
	rest,err := AppInit.ConnEs().Search().Query(BoolMustQuery). //bool查询
		SortBy(sortList...).	//排序
		From((search.Current-1)* search.Size).Size(search.Size).  //分页
		Index("books").Do(ctx)
	if err != nil {
		ctx.JSON(500,gin.H{"error":err})
	}else{
		ctx.JSON(200,gin.H{"result": ResultToBooks(rest),"metas":gin.H{"total":rest.TotalHits()}})
	}
}






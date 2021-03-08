package main

import (
	"context"
	"github.com/olivere/elastic/v7"
	"log"
	"strconv"
	"sync"
	"test/ES/AppInit"
	"test/ES/Model"
)

//
//func mainx(){
//
//	pagesize :=500	//
//	page := 1 	//页
//	wg := sync.WaitGroup{}	//
//	for {
//		//1、创建一个结构体切片，用来存储所有的数据查询结果
//		book_list := Model.BookSlice{}
//		//取500条数据
//		db := AppInit.GetDB().Order("book_id desc").
//			Limit(pagesize).Offset((page-1)*pagesize).Find(&book_list)
//		if db.Error != nil || len(book_list) == 0 {
//			break
//		}
//
//		wg.Add(1)
//		go func(){
//			defer wg.Done()
//			//2、创建Bulk(),是是批量插入/更新/删除文档的入口点。
//			ctx := context.Background()
//			bulk := AppInit.ConnEs().Bulk()
//			for _,book := range book_list{
//
//				//批量插入数据，返回一个新的bulkkindexrequest。默认操作类型为“index”
//				req:=elastic.NewBulkIndexRequest()
//				//elastic.NewBulkUpdateRequest()	//批量更新
//				//elastic.NewBulkDeleteRequest()	//批量删除
//
//				//3、向books索引中，写入数据，指定数据的 Id(ID) \ Doc(内容)
//				req.Index("books").Id(strconv.Itoa(book.BookID)).Doc(book)
//				//4、循环处理，每处理完一个，就再加入一个
//				bulk.Add(req)
//			}
//			//5、最后确认执行bulk批量插入
//			_,err := bulk.Do(ctx)
//			if err != nil {
//				log.Println("bulk插入数据失败",err)
//			}else{
//				fmt.Println("bulk插入数据成功")
//			}
//		}()
//		page = page+1	//分页+1
//	}
//	wg.Wait()
//}

func main(){


	//1.连接数据库拿取数据
	pageSize :=500
	page := 1
	wg := sync.WaitGroup{}

	for {
		books := Model.BookSlice{}

		db := AppInit.GetDB().Order("book_id desc").Limit(pageSize).
			Offset((page - 1) * pageSize).Find(&books)
		if db.Error != nil || len(books) == 0 {
			break
		}

		wg.Add(1)
		go func(){
			defer wg.Done()
			ctx := context.Background()
			bulk := AppInit.ConnEs().Bulk()

			for _,book := range books{
				req := elastic.NewBulkIndexRequest()
				req.Index("books").Id(strconv.Itoa(book.BookID)).Doc(book)
				bulk.Add(req)
			}
			_,err := bulk.Do(ctx)
			if err != nil {
				log.Println("bulk插入错误")
			}else{
				log.Println("bulk插入成功")
			}
		}()
		page = page+1
	}
	wg.Wait()
}
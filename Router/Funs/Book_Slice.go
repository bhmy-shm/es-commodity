package Funs

import (
	"elasticsearch/Model"
	"github.com/olivere/elastic/v7"
	"reflect"
)

func ResultToBooks(rsp *elastic.SearchResult) []*Model.Books{
	ret := []*Model.Books{}
	var t  *Model.Books
	for _,item := range rsp.Each(reflect.TypeOf(t)){
		ret = append(ret,item.(*Model.Books))
	}
	return ret
}

func ResultToSource(rest *elastic.SearchResult) []interface{}{
	ret := make([]interface{},0)

	for _,hit := range rest.Hits.Hits{
		ret = append(ret,hit.Source)
	}
	return ret
}

func ResultToFileds(rsp *elastic.SearchResult,key string) []interface{} {
	ret:=make([]interface{},0)
	for _,hit:=range rsp.Hits.Hits{
		ret=append(ret,hit.Fields[key].([]interface{})[0])
	}
	return ret
}

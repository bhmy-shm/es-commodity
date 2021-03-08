package AppInit

import (
	"github.com/olivere/elastic/v7"
	"log"
)

//初始化连接 ES
func ConnEs() *elastic.Client{

	client,err := elastic.NewClient(
		elastic.SetURL("http://192.168.168.4:9200"),
		elastic.SetSniff(false),
		)
	if err != nil {
		log.Println(err)
		return nil
	}
	return client
}
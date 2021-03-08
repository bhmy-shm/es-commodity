package AppInit

import (
	"log"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

var db  *gorm.DB
func init() {
	var err error
	db, err = gorm.Open("mysql",
		"shm:123.com@tcp(192.168.168.4:3306)/book?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
	}else{
		log.Println("数据库连接成功")
	}

	//db.LogMode(true)             //设置日志模式
	db.SingularTable(true)     //设置结构体表的复数形式
	db.DB().SetMaxOpenConns(10)  //空闲连接池中的最大连接数
	db.DB().SetMaxOpenConns(100) //设置数据库连接最大打开数
	//db.DB().SetConnMaxLifetime(time.Hour)  设置可宠用连接的最长时间
}

func GetDB() *gorm.DB{
	return db
}
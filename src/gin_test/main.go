package main
import (
	//这里讲db作为go/databases的一个别名，表示数据库连接池
	db "gin_test/databases"
	. "gin_test/router"
)

var dbCon = db.GetDbConnect()

func main() {
	//当整个程序完成之后关闭数据库连接
	defer dbCon.Close()
	router := InitRouter()
	router.Run(":8080")
}
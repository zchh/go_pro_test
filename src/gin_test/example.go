package main

import (
	"database/sql"
	"fmt"
	//"github.com/gin-gonic/gin"
	//"log"
	//"net/http"
	"time"

	//"github.com/golang/protobuf/protoc-gen-go/descriptor"
	//
	//"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	//"net/http"

	_ "github.com/go-sql-driver/mysql"
	//"io"
	//"os"
	"reflect"
)

// Booking contains binded and validated data.
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

type Person struct {
	uid        int    `json:"uid"`
	username string `json:"username"`
	department  string `json:"department"`
	created string `json:"created"`
}

type Personslice struct {
	 persons []Person
}


//var d = Driver{proto: "tcp", raddr: "127.0.0.1:3306"}
//func init() {
//	Register("SET NAMES utf8")
//	sql.Register("mymysql", &d)
//}





func bookableDate(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}






func main()  {

	db,_ := sql.Open("mysql", "root:123456@tcp(127.0.0.1:12330)/test?charset=utf8")
	//checkErr(err)

	fmt.Println(reflect.TypeOf(db))
	//fmt.Println(reflect.TypeOf(err))



	//插入数据
	//stmt, err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	////checkErr(err)
	//res, err := stmt.Exec("hhh", "产品部门", "2012-12-12")
	////checkErr(err)
	//id, err := res.LastInsertId()
	////checkErr(err)
	//fmt.Println(err)
	//fmt.Println(id)


	//删除数据
	//id := 4
	//stmt, err := db.Prepare("delete from userinfo where uid=?")
	//res, err := stmt.Exec(id)
	//affect, err := res.RowsAffected()
	//fmt.Println(affect)
	//fmt.Println(err)

	//更新数据
	//stmt, err := db.Prepare("update userinfo set username=? where uid=?")
	//res, err := stmt.Exec("astaxieupdate", 3)
	//affect, err := res.RowsAffected()
	//fmt.Println(err)
	//fmt.Println(affect)



	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	defer rows.Close()
	if err != nil{
		fmt.Println(err)
	}

	var single map[string]interface{}
	var data []map[string]interface{}
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		rows.Scan(&uid, &username, &department, &created)

		defer rows.Close()

		single = make(map[string]interface{})
		single["uid"] = uid
		single["username"] = username
		single["department"] = department
		single["created"] = created
		data = append(data, single)
	}


	//fmt.Println(persons)

	db.Close()

	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"persons": persons,
	//	})
	//})
	//r.Run()




	//var m map[string]interface{}
	//var s []map[string]interface{}
	//m = make(map[string]interface{})
	//m["username"] = "user1"
	//m["age"] = 18
	//m["sex"] = "man"
	//s = append(s, m)
	//m = make(map[string]interface{})
	//m["username"]="user2"
	//m["age"]=188
	//m["sex"]="male"
	//s=append(s,m)
	//data, err := json.Marshal(s)
	//if err != nil {
	//	fmt.Printf("json.marshal failed,err:", err)
	//	return
	//}
	//fmt.Printf("%s\n", string(data))






	//// Logging to a file.
	//gin.ForceConsoleColor()
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//
	//
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//		"data": data,
	//	})
	//})
	//r.Run() // listen and serve on 0.0.0.0:8080
}




package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	//contact mysql
	db, err := sql.Open("mysql", "root:password@tcp(locahost:3306)/test")
	if err != nil {
		log.Println("数据库连接失败: ", err)
		return
	}
	db.Ping()
	defer func() {
		db.Close()
	}()

	stmt, err := db.Prepare("delete from test where id = ?")
	if err != nil {
		log.Println("预处理失败: ", err)
		return
	}
	defer func() {
		stmt.Close()
	}()

	res,err:=stmt.Exec(1)
	if err!=nil{
		fmt.Println("执行SQL出现错误")
	}
	//受影响行数
	count,err:=res.RowsAffected()
	if err!=nil{
		fmt.Println("获取结果失败",err)
	}
	if count > 0{
		fmt.Println("删除成功")
	} else {
		fmt.Println("删除失败")
	}
}

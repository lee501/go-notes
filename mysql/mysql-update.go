package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/test")
	if err != nil {
		log.Println("mysql连接失败: ", err)
		return
	}
	db.Ping()
	//关闭连接
	defer func() {
		db.Close()
	}()

	//sql预处理，使用占位符
	stmt, err := db.Prepare("update test set name = ?, address = ? where id = ?")
	if err != nil {
		log.Println("sql预处理失败: ", err)
		return
	}

	defer func() {
		stmt.Close()
	}()

	//exec执行函数，对应占位符
	result, err := stmt.Exec("anne", "东城",  1)
	if err != nil {
		log.Println("sql执行失败: ", err)
		return
	}

	count, _ := result.RowsAffected()
	if count > 0{
		fmt.Println("数据修改成功")
	} else {
		fmt.Println("数据更新失败")
	}
}

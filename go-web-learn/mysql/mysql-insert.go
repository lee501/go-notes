package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//打开连接
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/test")
	db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}

	defer func() {
		db.Close()
	}()

	//预处理sql
	stmt, err := db.Prepare("insert into test values(default, ?, ?)")
	if err != nil {
		fmt.Println("预处理失败")
		return
	}

	defer func() {
		stmt.Close()
	}()

	r, err := stmt.Exec("lee", "北京")
	if err != nil {
		fmt.Println("sql执行失败")
		return
	}
	count, err := r.RowsAffected()
	if err != nil {
		fmt.Println("数据插入失败")
		return
	}

	if count > 0 {
		fmt.Println("数据插入成功")
	} else {
		fmt.Println("失败")
	}
}

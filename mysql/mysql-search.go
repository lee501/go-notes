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
		log.Println("数据库连接失败: ", err)
		return
	}
	db.Ping()
	defer func() {
		db.Close()
	}()

	stmt, err := db.Prepare("select name from test where id = ?")
	if err != nil {
		log.Println("预处理失败: ", err)
		return
	}
	defer func() {
		stmt.Close()
	}()

	rows, err := stmt.Query(1)
	for rows.Next() {
		var name string
		fmt.Println(rows.Scan(&name))
	}
}

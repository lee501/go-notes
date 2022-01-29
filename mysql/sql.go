package main

import (
"database/sql"
"fmt"
_ "github.com/go-sql-driver/mysql"
)

type EgoResult struct {
	Status int
	Data   interface{}
	Msg    string
}

var (
	db 	 *sql.DB
	stmt *sql.Stmt
	rows *sql.Rows
)

func openConn() (err error) {
	db, err = sql.Open("mysql","root:password@tcp(localhost:3306)/test")
	if err != nil {
		fmt.Println("mysql连接失败", err)
		return
	}
	return
}

func CloseConn() {
	if rows != nil {
		rows.Close()
	}
	if stmt != nil {
		stmt.Close()
	}
	if db != nil {
		db.Close()
	}
}
//执行新增 删除 修改
func Dml(sql string, args ...interface{}) (int64, error) {
	err := openConn()
	defer func() {
		CloseConn()
	}()
	if err != nil {
		fmt.Println("执行DML报错，打开连接失败")
		return 0, err
	}
	stmt, err = db.Prepare(sql)
	if err != nil {
		fmt.Println("执行DML报错，预处理失败")
		return 0, err
	}
	result, err := stmt.Exec(args...)
	if err != nil {
		fmt.Println("执行DML报错，数据执行失败")
		return 0, err
	}
	count, err := result.RowsAffected()
	return  count, err
}

//执行查询
func Dql(sql string,args ... interface{}) (*sql.Rows,error){
	err:= openConn()
	if err!=nil{
		fmt.Println("执行DQL出现错误,打开连接失败")
		return nil,err
	}
	//此处是等号
	stmt,err=db.Prepare(sql)
	if err!=nil{
		fmt.Println("执行DQL出现错误,预处理实现")
		return nil,err
	}
	//此处参数是切片
	rows,err=stmt.Query(args...)
	if err!=nil{
		fmt.Println("执行DQL出现错误,执行错误")
		return nil,err
	}
	//此处没有关闭,调用此函数要记得关闭连接
	return rows,nil
}
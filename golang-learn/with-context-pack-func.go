package main

import "sync"

//函数重复的部分，封装到一个函数withContext中，函数作为参数，在调用者的上下文来调用withContext

func withContext(fn func()) {
	mu := new(sync.Mutex)
	mu.Lock()
	defer mu.Unlock()
	fn()
}

func bar() {
	withContext(func() {
		// do some action for bar
	})
}

type DB struct {

}
//数据库链接
func withDBContext(fn func(db *DB) error) error{
	// 从连接池获取一个数据库连接
	dbConn := NewDB()
	return fn(dbConn)
}

func NewDB() *DB{
	return &DB{}
}